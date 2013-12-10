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
import "bytes"
import "sync"

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
	content string
	mu sync.Mutex
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
	Title, Author, Time string
	Text template.HTML
}

type mainTemplate struct {
	Title string
	Entries []entryTemplate
}

type tagTemplate struct {
	Tags []string
}

// unix timestamp to string
func convertTime(inttime int) string {
	const layout = "Jan 2, 2006 at 3:04pm (MST)"
	t := time.Unix(int64(inttime), 0)
	return t.Format(layout)
}

// read one blog
func (b *Blog) readEntry(entry *entryTemplate, idx int) {
	entry.Title = b.indices[idx].Title
	entry.Time = convertTime(b.indices[idx].Time)
	buf, err := ioutil.ReadFile(b.indices[idx].Text)
	if err != nil {
		fmt.Println(err)
		return
	}
	entry.Text = template.HTML(buf)
}

// read all blogs from time start
func (b *Blog) readBlog(data *mainTemplate, start int) {
	data.Title = b.Title

	i := len(b.indices)
	data.Entries = make([]entryTemplate, i)
	for j := 0; j < i; j++ {
		b.readEntry(&data.Entries[j], j)
	}
}

// create blog cache
func (b *Blog) blogCache() string {
//	if (b.content != "") {
//		return b.content
//	}

	b.mu.Lock()
	defer b.mu.Unlock()

	t, err := template.ParseFiles("../data/blog/template/main2.template")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	var data mainTemplate
	b.readBlog(&data, -1)

	output := new(bytes.Buffer)
	t.Execute(output, data)
	b.content, _ = output.ReadString(0)
	return b.content
}

// call by http server
func (b *Blog) Serve(w io.Writer, req *http.Request) error {
	io.WriteString(w, b.blogCache())
	return nil
}

func NewBlog(title string, indexPath string) *Blog {
	b := new(Blog)
	b.Title = title
	b.IndexPath = indexPath
	if err := b.readIndex(); err != nil {
		return nil
	}
	return b
}

func TestBlog() {
	b := NewBlog("TestBlog", "../data/blog/index.json")
	io.WriteString(os.Stdout, b.blogCache())
}