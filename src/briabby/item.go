package briabby

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"sync"
)

type HatItem struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	Promotion  string   `json:"promotion"`
	ImageSmall string   `json:"smallimage"`
	ImageBig   string   `json:"bigimage"`
	Desc       string   `json:"desc"`
	Price      []string `json:"price"`
	PaypalBtn  string   `json:"paypalbtn"`
	Category   string   `json:"category"`
}

var hatitemarray []*HatItem
var hatitemmap map[int]*HatItem
var maxhatitemid int
var itemlock sync.Mutex
var itemfile = "../data/babegarden/item.json"

type ItemArray []*HatItem

func (ia ItemArray) Len() int      { return len(ia) }
func (ia ItemArray) Swap(i, j int) { ia[i], ia[j] = ia[j], ia[i] }
func (ia ItemArray) Less(i, j int) bool {
	a := len(ia[i].Promotion)
	b := len(ia[j].Promotion)
	if a == b {
		return ia[i].ID < ia[i].ID
	}
	return a > b
}

func GetNextItemID() int {
	id := maxhatitemid
	maxhatitemid++
	return id
}

func FindItemByID(id int) *HatItem {
	if item, ok := hatitemmap[id]; ok {
		return item
	}

	return nil
}

func FindItemIdxByID(id int) int {
	for i := range hatitemarray {
		if hatitemarray[i].ID == id {
			return i
		}
	}

	return -1
}

func SaveItem(item *HatItem) error {
	itemlock.Lock()
	defer itemlock.Unlock()

	item.ID = GetNextItemID()
	hatitemarray = append(hatitemarray, item)
	sort.Sort(ItemArray(hatitemarray))
	hatitemmap[item.ID] = item
	err := FlushFile()
	return err
}

func EditItem(item *HatItem) error {
	itemlock.Lock()
	defer itemlock.Unlock()

	if i := FindItemIdxByID(item.ID); i != -1 {
		hatitemarray[i] = item
	} else {
		return errors.New("not found")
	}
	hatitemmap[item.ID] = item
	return nil
}

func DelItem(id int) error {
	itemlock.Lock()
	defer itemlock.Unlock()

	if i := FindItemIdxByID(id); i != -1 {
		fmt.Printf("delete item %d in idx %d\n", hatitemmap[i].ID, i)
		hatitemarray = append(hatitemarray[:i], hatitemarray[i+1:]...)
		delete(hatitemmap, id)
	}
	return nil
}

func FlushFile() error {
	b, err := json.Marshal(hatitemarray)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(itemfile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	var out bytes.Buffer
	json.Indent(&out, b, "", "\t")
	_, err = f.Write(out.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func InitItemsFromJson() error {
	hatitemmap = make(map[int]*HatItem)
	f, err := os.OpenFile(itemfile, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &hatitemarray)
	if err != nil {
		return err
	}
	for i := range hatitemarray {
		hatitemmap[hatitemarray[i].ID] = hatitemarray[i]
		if hatitemarray[i].ID >= maxhatitemid {
			maxhatitemid = hatitemarray[i].ID + 1
		}
	}
	if maxhatitemid == 0 {
		maxhatitemid = 1
	}
	fmt.Println(hatitemarray)
	sort.Sort(ItemArray(hatitemarray))
	return nil
}
