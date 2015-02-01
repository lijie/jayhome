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
	http.Handle("/images2/", http.StripPrefix("/images2/", http.FileServer(http.Dir(*rootDir + "/images2/"))))
	http.Handle("/template/", http.StripPrefix("/template/", http.FileServer(http.Dir(*rootDir + "/template/"))))
	http.Handle("/children/", http.StripPrefix("/children/", http.FileServer(http.Dir(*rootDir + "/children/"))))
	http.Handle("/admin/", http.StripPrefix("/admin/", http.FileServer(http.Dir(*rootDir + "/admin/"))))
	http.Handle("/tmp/", http.StripPrefix("/tmp/", http.FileServer(http.Dir("../tmp/"))))

	briabby.InitBriabby("")
	
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		log.Fatal("fatal %p", err)
	}
}
