package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Page struct {
	PageId int    `json:"pageId"`
	Title  string `json:"title"`
}

type SearchInfo struct {
	Totalhits int `json:"totalhits"`
}

type Search struct {
	Title   string `json:"title"`
	Snippet string `json:"snippet"`
}

type Query struct {
	Pages      map[string]Page `json:"pages"`
	SearchInfo SearchInfo      `json:"searchinfo"`
	Search     []Search        `json:"search"`
}

type WikipediaResponse struct {
	BatchComplete string `json:"batchcomplete"`
	Query         Query  `json:"query"`
}

func main() {
	// URL作成
	url := url.URL{}
	url.Scheme = "http"
	url.Host = "ja.wikipedia.org"
	url.Path = "w/api.php"
	query := url.Query()
	query.Set("action", "query")
	query.Set("list", "search")
	query.Set("srsearch", "じゃがいも")
	query.Set("format", "json")
	url.RawQuery = query.Encode()
	fmt.Println(url.String())

	// HTTPリクエスト
	resp, err := http.Get(url.String())
	if err != nil {
		fmt.Println("Error!")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	wikipediaResponse := new(WikipediaResponse)
	err = json.Unmarshal(body, wikipediaResponse)

	fmt.Printf("%+v\n", wikipediaResponse.Query.Search)
}
