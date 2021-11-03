package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Sendurl(term string, limits int) {
	if limits < 0 || limits > 200 {
		panic("limit must be smaller than 200")
	}
	limit := strconv.Itoa(limits)
	url := "https://itunes.apple.com/search?country=tw&media=movie&limit=" + limit + "&term=" + term
	var api_json []byte
	if response, err := http.Get(url); err != nil {
		fmt.Printf("request error: %v\n", err)
	}else if api_json, err = ioutil.ReadAll(response.Body); err != nil {
		fmt.Printf("read body error: %v\n", err)
	}
	Unmarshalapijson(api_json)
}