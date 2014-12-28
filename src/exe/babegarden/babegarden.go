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
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir(*rootDir + "/images/"))))
	http.Handle("/template/", http.StripPrefix("/template/", http.FileServer(http.Dir(*rootDir + "/template/"))))
	http.Handle("/children/", http.StripPrefix("/children/", http.FileServer(http.Dir(*rootDir + "/children/"))))

	briabby.InitBriabby("")
	
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		log.Fatal("fatal %p", err)
	}
}
