package blog

import "encoding/json"
import "os"
import "io/ioutil"
import "fmt"
import "sort"
import "html/template"
import "net/http"
import "io"
import "time"

type Index struct {
	Title string
	Tag string
	Text string
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

type entryTemplate struct {
	Title, Author, Time, Text string
}

type mainTemplate struct {
	Title string
	Entries []entryTemplate
}

func convertTime(inttime int) string {
	const layout = "Jan 2, 2006 at 3:04pm (MST)"
	t := time.Unix(int64(inttime), 0)
	return t.Format(layout)
}

func (b *Blog) readEntry(entry *entryTemplate, idx int) {
	entry.Title = b.indices[idx].Title
	entry.Time = convertTime(b.indices[idx].Time)
	buf, err := ioutil.ReadFile(b.indices[idx].Text)
	if err != nil {
		fmt.Println(err)
		return
	}
	entry.Text = string(buf)
}

func (b *Blog) readBlog(data *mainTemplate, start int) {
	b.readIndex()
	data.Title = b.Title

	i := len(b.indices)
	data.Entries = make([]entryTemplate, i)
	for j := 0; j < i; j++ {
		b.readEntry(&data.Entries[j], j)
	}
}

func (b *Blog) showMain(w io.Writer) error {
	t, err := template.ParseFiles("../data/blog/template/main.template")
	if err != nil {
		fmt.Println(err)
		return err
	}
	var data mainTemplate
	b.readBlog(&data, -1)
	t.Execute(w, data)
	return nil
}

func (b *Blog) Serve(w http.ResponseWriter, req *http.Request) error {
	return b.showMain(w)
}

func NewBlog(title string, indexPath string) *Blog {
	b := &Blog{title, indexPath, nil}
	if err := b.readIndex(); err != nil {
		return nil
	}
	return b
}

func TestBlog() {
	b := NewBlog("TestBlog", "../data/blog/index.json")
	b.showMain(os.Stderr)
}