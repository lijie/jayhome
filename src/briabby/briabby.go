package briabby

import (
	"fmt"
	"net/http"
	"text/template"
)

type HatDataSet struct {
	CurrentPage int
	MaxPage     int
	ItemList    []ProtoItem
}

func HandleHat(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../data/babegarden/template/hat2.html")
	if err != nil {
		return
	}

	var items []ProtoItem
	if items, err = store.FindItemByCat("hat"); err != nil {
		fmt.Println(err)
		return
	}

	ds := &HatDataSet{
		ItemList:    items,
		CurrentPage: 1,
	}

	if err = t.Execute(w, ds); err != nil {
		return
	}
}

type IndexDataSet struct {
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../data/babegarden/template/index.html")
	if err != nil {
		return
	}

	ds := &IndexDataSet{}

	if err = t.Execute(w, ds); err != nil {
		return
	}
}

var store *Store

func InitBriabby(prefix string) {
	var err error
	store, err = NewStore("162.243.132.159:27017")
	if err != nil {
		fmt.Printf("connect database error %v\n", err)
	} else {
		fmt.Printf("connect database ok\n")
	}

	http.HandleFunc(prefix+"/hat", HandleHat)
	http.HandleFunc(prefix+"/", HandleIndex)
	InitAdmin(prefix)
}
