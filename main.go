package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type WordPair struct {
	Word      string
	Occurence int
}

type WordPairList []WordPair

func main() {

	data, err := ioutil.ReadFile("text_file.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	text := string(data)
	a := strings.NewReader(text)

	p := WordPairList{}

	resp, err := http.Post("http://localhost:8080/data", "text/html", a)

	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(body, &p)
	fmt.Println(p)

}
