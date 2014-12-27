package main

import "net/http"
import "io"
import "log"
import "blog"
import "reader"
import "fmt"
import "flag"
import "sync"

var bb *blog.Blog
var blogMutex sync.Mutex

var rr *reader.Reader
var readerMutex sync.Mutex

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func BlogServer(w http.ResponseWriter, req *http.Request) {
	if bb == nil {
		blogMutex.Lock()
		if bb == nil {
			bb = blog.NewBlog("LI JIE", "../data/blog/index.json")
		}
		blogMutex.Unlock()
	}
	bb.Serve(w, req)
}

func ReaderServer(w http.ResponseWriter, req *http.Request) {
	if rr == nil {
		readerMutex.Lock()
		if rr == nil {
			rr = reader.NewReader()
		}
		readerMutex.Unlock()
	}
	rr.Serve(w, req)
}

var port = flag.Int("port", 8080, "default port")
var rootDir = flag.String("rootdir", "../data/static", "default root dir")

func main() {
	flag.Parse()

	// run for jayhome
	// serve static under an alternate URL
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(*rootDir))))
	http.HandleFunc("/b/", BlogServer)
	http.HandleFunc("/r/", ReaderServer)
	http.HandleFunc("/hello", HelloServer)

	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		log.Fatal("fatal %p", err)
	}
}
