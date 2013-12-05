package main

import "encoding/json"
import _ "os"
import "io/ioutil"
import "fmt"

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

func main() {
	buf, _ := ioutil.ReadFile("index.json")
	fmt.Println(string(buf))

	var idx []Index
	err := json.Unmarshal(buf, &idx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(idx)
}
