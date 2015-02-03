package briabby

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"
)

// FileInfo describes a file that has been uploaded.
type FileInfo struct {
	Key          string `json:"-"`
	Url          string `json:"url,omitempty"`
	ThumbnailUrl string `json:"thumbnail_url,omitempty"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Size         int64  `json:"size"`
	Error        string `json:"error,omitempty"`
	DeleteUrl    string `json:"delete_url,omitempty"`
	DeleteType   string `json:"delete_type,omitempty"`
}

type FileSave struct {
	reader   *multipart.Reader
	fileName string
	currfile *os.File
}

var errorJsonResp = "{\"result\":\"error\"}"
var okJsonResp = "{\"result\":\"ok\"}"

func (fs *FileSave) Save() *FileInfo {
	name := fmt.Sprintf("%d", time.Now().UnixNano()) + ".png"
	info := &FileInfo{
		Type: "unknown",
		Url:  "/tmp/" + name,
	}

	var n int64

	for {
		part, err := fs.reader.NextPart()
		if err == io.EOF {
			break
		}

		if part.FormName() == "" {
			break
		}
		if part.FileName() == "" {
			break
		}
		// fs.fileName = part.FileName()
		fs.fileName = name
		info.Name = part.FileName()
		if fs.currfile == nil {
			fs.currfile, _ = os.OpenFile("../tmp/"+fs.fileName, os.O_RDWR|os.O_CREATE, 0666)
			defer fs.currfile.Close()
		}
		n, _ = io.Copy(fs.currfile, part)
		info.Size += n
	}

	return info
}

func NewFileSave(r *multipart.Reader) *FileSave {
	return &FileSave{
		reader: r,
	}
}

func HandleAdminLogin(w http.ResponseWriter, r *http.Request) {
}

func fnHandleSaveItem(w http.ResponseWriter, r *http.Request) {
	var item *HatItem

	id := r.FormValue("item_id")
	if len(id) == 0 || id == "0" {
		item = &HatItem{}
	} else {
		intid, _ := strconv.Atoi(id)
		item = FindItemByID(intid)
		if item == nil {
			io.WriteString(w, errorJsonResp)
			return
		}

		itemlock.Lock()
		defer itemlock.Unlock()
		item.ID = intid
	}

	item.Name = r.FormValue("item_name")
	item.Desc = r.FormValue("item_desc")
	item.PaypalBtn = r.FormValue("item_paypal")
	item.ImageSmall = r.FormValue("item_small_image_url")
	item.ImageBig = r.FormValue("item_big_image_url")
	item.Promotion = r.FormValue("item_promotion")
	item.Category = r.FormValue("item_category")
	item.Price = strings.Split(r.FormValue("item_price"), ",")
	if item.ID == 0 {
		SaveItem(item)
	} else {
		FlushFile()
	}
	io.WriteString(w, okJsonResp)
}

func fnHandleDelItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	if err == nil {
		DelItem(id)
	}
	fnHandleListItem(w, r)
}

func fnHandleEditItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		return
	}
	item := FindItemByID(id)
	if item == nil {
		return
	}
	t, err := template.ParseFiles("../data/babegarden/template/admin_item.html")
	if err != nil {
		return
	}
	fmt.Println(item)
	if err = t.Execute(w, item); err != nil {
		return
	}
}

func fnHandleListItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("xxx")
	t, err := template.ParseFiles("../data/babegarden/template/admin_item_list.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	ds := &HatDataSet{
		ItemList: hatitemarray,
	}

	itemlock.Lock()
	defer itemlock.Unlock()

	if err = t.Execute(w, ds); err != nil {
		return
	}

}

func HandleAdminItem(w http.ResponseWriter, r *http.Request) {
	fn := r.FormValue("fn")
	defer r.Body.Close()
	if len(fn) == 0 {
		t, err := template.ParseFiles("../data/babegarden/template/admin_item.html")
		if err != nil {
			return
		}

		item := &HatItem{}
		if err = t.Execute(w, item); err != nil {
			return
		}
	}

	if fn == "add" {
		fnHandleSaveItem(w, r)
		return
	}

	if fn == "edit" {
		fnHandleEditItem(w, r)
	}

	if fn == "del" {
		fnHandleDelItem(w, r)
		return
	}

	if fn == "list" {
		fnHandleListItem(w, r)
		return
	}
}

func HandleAdminUpload(w http.ResponseWriter, r *http.Request) {
	mr, err := r.MultipartReader()
	if err != nil {
		fmt.Printf("multipart reader err %v\n", err)
		return
	}
	info := NewFileSave(mr).Save()

	out, _ := json.Marshal(info)
	// fmt.Println(string(out))
	w.Write(out)
}

func InitAdmin(prefix string) {
	http.HandleFunc(prefix+"/admin/login", HandleAdminLogin)
	http.HandleFunc(prefix+"/admin/item", HandleAdminItem)
	http.HandleFunc(prefix+"/admin/upload", HandleAdminUpload)
}
