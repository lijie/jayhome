package briabby

import (
	"net/http"
	"html/template"
	"fmt"
)

type HatItem struct {
	Title string
	Image string
	Desc  string
}

type HatDataSet struct {
	ItemList []HatItem
}

func HandleHat(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandleHat")
	t, err := template.ParseFiles("../data/briabby/hat.html")
	if err != nil {
		return
	}

	ds := &HatDataSet{
		ItemList: make([]HatItem, 6),
	}

	for i := range ds.ItemList {
		item := &ds.ItemList[i]
		item.Title = "Awesome Hat"
		item.Image = "/briabby/aviator_hat.jpg"
		item.Desc = "Very cute hat for children!"
	}

	if err = t.Execute(w, ds); err != nil {
		return
	}
}

func InitBriabby(prefix string) {
	http.HandleFunc("/briabby/hat", HandleHat)
}
