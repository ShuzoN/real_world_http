package main

import (
	"log"
	"net/http"
	"net/url"
)

func main() {
	// request
	values := url.Values{
		"test": {"value"},
	}

	// response
	resp, err := http.PostForm("http://localhost:18888", values)
	//resp, err := http.Head("http://localhost:18888")
	if err != nil {
		panic(err)
	}
	log.Println("Status", resp.Status)
}
