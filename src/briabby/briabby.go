package briabby

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"encoding/csv"
)

type HatItem struct {
	ID         int
	Name       string
	ImageSmall string
	ImageBig   string
	Desc       string
	Price      [7]float64
}

type HatData struct {
	ItemList []*HatItem
}

type HatDataSet struct {
	ItemList []*HatItem
}

func (hd *HatData) getPage(pagenum int) []*HatItem {
	count := 9
	pagenum--
	if pagenum*count >= len(hd.ItemList) {
		return nil
	}
	if len(hd.ItemList) - pagenum*count > 9 {
		return hd.ItemList[pagenum*count:pagenum*count+9]
	} else {
		return hd.ItemList[pagenum*count:]
	}
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
			Name:       record[1],
			Desc:       record[2],
			ImageSmall: record[3],
			ImageBig:   record[4],
		}
		item.ID, _ = strconv.Atoi(record[0])

		for i := range item.Price {
			item.Price[i], _ = strconv.ParseFloat(record[i+4], 32)
		}
		hd.ItemList = append(hd.ItemList, item)
		fmt.Println(item)
	}
	return nil
}

func HandleHat(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandleHat")
	t, err := template.ParseFiles("../data/briabby/hat.html")
	if err != nil {
		return
	}

	page := 1
	pagestr := r.FormValue("page")
	if len(pagestr) > 0 {
		page, _ = strconv.Atoi(pagestr)
	}

	ds := &HatDataSet{
		ItemList: hatdata.getPage(page),
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
