package main

import "net/http"
import "io"
import "log"
import "blog"
import "fmt"
import "flag"
import "sync"

var bb *blog.Blog
var blogMutex sync.Mutex

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func BlogServer(w http.ResponseWriter, req *http.Request) {
	if bb == nil {
		blogMutex.Lock()
		if bb == nil {
			bb = blog.NewBlog("TestBlog", "../data/blog/index.json")
		}
		blogMutex.Unlock()
	}
	bb.Serve(w, req)
}

var port = flag.Int("port", 80, "default port")

func main() {
	flag.Parse()
	// serve static under an alternate URL
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("../data/static"))))
	http.HandleFunc("/b", BlogServer)
	http.HandleFunc("/hello", HelloServer)

	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		log.Fatal("fatal %p", err)
	}
}
