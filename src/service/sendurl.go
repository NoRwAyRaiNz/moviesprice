package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Sendurl(term string) {
	url := "https://itunes.apple.com/search?country=tw&media=movie&limit=1&term=" + term
	var api_json []byte
	if response, err := http.Get(url); err != nil {
		fmt.Printf("request error: %v\n", err)
	}else if api_json, err = ioutil.ReadAll(response.Body); err != nil {
		fmt.Printf("read body error: %v\n", err)
	}
	Unmarshalapijson(api_json)
}