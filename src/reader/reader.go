package reader

import "fmt"
import _ "os"
import "encoding/xml"
import "encoding/json"
import "io/ioutil"
import "io"
import "sort"
import "html/template"
import "time"
import "strings"
import "strconv"
import "net/http"

type Item struct {
	Title string `xml:"title"`
	Link string `xml:"link"`
	Description template.HTML `xml:"description"`
}

type Channel struct {
	XMLName xml.Name `xml:"channel"`
	ItemList []Item `xml:"item"`
}

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel `xml:"channel"`
}

type DataSet struct {
	ItemList []Item
	Indices []Index
}

type Index struct {
	Title, Tag, Url string
	Id int
	content []byte
	timeout time.Time
}

type Reader struct {
	IndexPath string
	indices []Index
}

type byId []Index
func (b byId) Len() int { return len(b) }
func (b byId) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b byId) Less(i, j int) bool { return b[i].Id < b[j].Id }

func (r *Reader) readIndex() error {
	buf, _ := ioutil.ReadFile(r.IndexPath)
	// fmt.Println(string(buf))

	err := json.Unmarshal(buf, &r.indices)
	if err != nil {
		return err
	}
	fmt.Println(r.indices)
	sort.Sort(byId(r.indices))
	return nil
}

func (r *Reader) fetchContent(idx *Index) {
	if time.Now().Before(idx.timeout) {
		fmt.Printf("no need fetch\n")
		return
	}

	// read content
	resp, err := http.Get(idx.Url)
	if err != nil {
		fmt.Printf("Get %s, err %v\n", idx.Url, err)
		return
	}

	defer resp.Body.Close()
	idx.content, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ReadAll %s, err %v\n", idx.Url, err)
		return
	}
	idx.timeout = time.Now().Add(time.Second * 3600)
}

func (r *Reader) ReadRSS(w io.Writer, id int) {
	var v Rss
	idx := &r.indices[id]

	r.fetchContent(idx)
	fmt.Printf("buf size %d\n", len(idx.content))
	xml.Unmarshal(idx.content, &v)
	// fmt.Println(v)
	var data DataSet
	data.ItemList = v.Channel.ItemList
	data.Indices = r.indices
	r.render(w, &data)
}

func (r *Reader) Serve(w io.Writer, req *http.Request) {
	sub := strings.Split(req.RequestURI, "/")
	if len(sub) >= 3 {
		if stamp, err := strconv.Atoi(sub[2]); err == nil {
			r.ReadRSS(w, stamp)
			return
		}
	}
	r.ReadRSS(w, 0)
}

func (r *Reader) render(w io.Writer, v *DataSet) {
	t, err := template.ParseFiles("../data/reader/template/main2.template")
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Execute(w, v)
}

func NewReader() *Reader {
	r := &Reader{}
	r.IndexPath = "../data/reader/index.json"
	r.readIndex()
	return r
}
