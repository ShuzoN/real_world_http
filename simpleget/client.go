package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	// request
	values := url.Values{
		"query": {"hello world"},
	}

	// response
	resp, err := http.Get("http://localhost:18888" + "?" + values.Encode())
	//resp, err := http.Head("http://localhost:18888")
	if err != nil {
		panic(err)
	}
	// main関数終了後に呼び出し
	// panicした場合も実行される
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println("Status", resp.Status)
	log.Println("Header", resp.Header)
	log.Println("Body: ", string(body))
	log.Println("Content-Lnegth:", resp.Header.Get("Content-Length"))
}
