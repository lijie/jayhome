package briabby

import (
	"fmt"
	"text/template"
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
	PaypalBtn  string
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
	PaypalBtn string
}

const paypalbtn = `
<form target="paypal" action="https://www.paypal.com/cgi-bin/webscr" method="post">
<input type="hidden" name="cmd" value="_s-xclick">
<input type="hidden" name="hosted_button_id" value="JCS44R3JUVXQA">
<table>
<tr><td><input type="hidden" name="on0" value="age"></td></tr><tr><td><select name="os0">
<option value="6-12MONTH">6-12MONTH $0.15 USD</option>
<option value="12-24MONTH">12-24MONTH $0.16 USD</option>
<option value="2-3YEAR">2-3YEAR $0.17 USD</option>
</select> </td></tr>
</table>
<input type="hidden" name="currency_code" value="USD">
<input type="image" src="https://www.paypalobjects.com/en_US/i/btn/btn_cart_LG.gif" border="0" name="submit" alt="PayPal - The safer, easier way to pay online!">
<img alt="" border="0" src="https://www.paypalobjects.com/en_US/i/scr/pixel.gif" width="1" height="1">
</form>`

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
			PaypalBtn: paypalbtn,
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
	t, err := template.ParseFiles("../data/babegarden/hat.html")
	if err != nil {
		return
	}

	ds := &HatDataSet{
		ItemList: hatdata.ItemList,
		PromotionList: hatdata.PromotionList,
		CurrentPage: 1,
	}

	if err = t.Execute(w, ds); err != nil {
		return
	}
}

var hatdata HatData
func InitBriabby(prefix string) {
	hatdata.initFromCSV("../data/babegarden/children/children_hat_list.csv")
	fmt.Println(hatdata)
	http.HandleFunc(prefix + "/hat", HandleHat)
}
