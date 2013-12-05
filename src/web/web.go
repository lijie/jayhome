package main

import "net/http"
import "io"
import "log"

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	// serve static under an alternate URL
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("static"))))
	// http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("fatal %p", err)
	}
}
