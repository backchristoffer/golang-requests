package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	request, error := http.Get("https://httpbin.org/get")
	if error != nil {
		log.Fatalln(error)
	}

	defer request.Body.Close()

	body, error := ioutil.ReadAll((request.Body))
	if error != nil {
		log.Fatalln(error)
	}

	log.Println(string(body))
}
