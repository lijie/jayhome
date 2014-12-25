package briabby

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"encoding/csv"
)

const (
	CSVHatID = iota
	CSVHatPromotion
	CSVHatName
	CSVHatDesc
	CSVHatSmallImage
	CSVHatBigImage
	CSVHatPrice
)

type HatItem struct {
	ID         int
	Name       string
	Promotion  int
	ImageSmall string
	ImageBig   string
	Desc       string
	Price      [7]float64
}

type HatData struct {
	ItemList []*HatItem
	PromotionList []*HatItem
}

type HatDataSet struct {
	CurrentPage int
	MaxPage int
	ItemList []*HatItem
	PromotionList []*HatItem
}

func (hd *HatData) getItemList(pagenum int) []*HatItem {
	return hd.ItemList
}

func (hd *HatData) getPromotion() []*HatItem {
	return hd.PromotionList
}

func (hd *HatData) initFromCSV(path string) error {
	f, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	r := csv.NewReader(f)
	// r.Comment = []rune("#")

	var record []string
	for {
		record, err = r.Read()
		if err != nil {
			break
		}
		item := &HatItem{
			Name:       record[CSVHatName],
			Desc:       record[CSVHatDesc],
			ImageSmall: record[CSVHatSmallImage],
			ImageBig:   record[CSVHatBigImage],
		}
		item.Promotion, _ = strconv.Atoi(record[CSVHatPromotion])
		item.ID, _ = strconv.Atoi(record[CSVHatID])

		for i := range item.Price {
			item.Price[i], _ = strconv.ParseFloat(record[i+CSVHatPrice], 32)
		}
		if item.Promotion != 0 {
			hd.PromotionList = append(hd.PromotionList, item)
		} else {
			hd.ItemList = append(hd.ItemList, item)
		}
		// fmt.Println(item)
	}
	return nil
}

func HandleHat(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandleHat")
	t, err := template.ParseFiles("../data/briabby/hat.html")
	if err != nil {
		return
	}

	ds := &HatDataSet{
		ItemList: hatdata.ItemList,
		PromotionList: hatdata.PromotionList,
		CurrentPage: 1,
		MaxPage: 1,
	}

	if err = t.Execute(w, ds); err != nil {
		return
	}
}

var hatdata HatData
func InitBriabby(prefix string) {
	hatdata.initFromCSV("../data/briabby/children/children_hat_list.csv")
	fmt.Println(hatdata)
	http.HandleFunc("/briabby/hat", HandleHat)
}
