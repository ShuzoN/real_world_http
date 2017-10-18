package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// request
	file, err := os.Open("../simpleget/client.go")
	if err != nil {
		panic(err)
	}

	// response
	resp, err := http.Post("http://localhost:18888", "text/plain", file)
	if err != nil {
		panic(err)
	}
	log.Println("Status", resp.Status)
}
