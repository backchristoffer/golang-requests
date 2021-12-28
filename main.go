package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	request, error := http.Get("http://date.jsonest.com/ll")
	if error != nil {
		log.Fatal("Was not able to get request", error)
	}
	defer request.Body.Close()
	response, error := io.ReadAll(request.Body)

	fmt.Printf(string(response))
}
