package main

import "net/http"
import "io"
import "log"
import "blog"
import "fmt"
import "flag"

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func BlogServer(w http.ResponseWriter, req *http.Request) {
	b := blog.NewBlog("TestBlog", "../data/blog/index.json")
	fmt.Println(b)
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
