package blog

import "encoding/json"
import "os"
import "io/ioutil"
import "fmt"
import "sort"
import "html/template"
import "net/http"

type Index struct {
	Title string
	Tag string
	Md string
	Time int
}

type Blog struct {
	Title string
	IndexPath string

	indices []Index
}

type byTime []Index
func (b byTime) Len() int { return len(b) }
func (b byTime) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b byTime) Less(i, j int) bool { return b[i].Time > b[j].Time }

func (b *Blog) readIndex() error {
	buf, _ := ioutil.ReadFile(b.IndexPath)
	fmt.Println(string(buf))

	err := json.Unmarshal(buf, &b.indices)
	if err != nil {
		return err
	}
	fmt.Println(b.indices)
	sort.Sort(byTime(b.indices))
	return nil
}

type mainTemplate struct {
	Title, Entries string
}

type entryTemplate struct {
	Title, Author, Time, Text string
}

func (b *Blog) showMain() error {
	t, err := template.ParseFiles("../data/blog/temlate/main.template")
	if err != nil {
		fmt.Println(err)
		return err
	}
	var data mainTemplate
	data.Title = "TestBlog"
	data.Entries = "TestEntries"
	t.Execute(os.Stdout, data)
	return nil
}

func (b *Blog) Serve(w http.ResponseWriter, req *http.Request) error {
	return b.showMain()
}

func NewBlog(title string, indexPath string) *Blog {
	b := &Blog{title, indexPath, nil}
	if err := b.readIndex(); err != nil {
		return nil
	}
	return b
}

