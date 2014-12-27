package main

import "net/http"
import "log"
import "fmt"
import "flag"
import "briabby"

var port = flag.Int("port", 8083, "default port")
var rootDir = flag.String("rootdir", "../data/babegarden", "default root dir")

func main() {
	flag.Parse()

	// serve static under an alternate URL
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(*rootDir))))

	briabby.InitBriabby("")
	
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		log.Fatal("fatal %p", err)
	}
}
