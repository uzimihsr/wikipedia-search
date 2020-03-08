package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

// JSONをパースするための構造体を定義
type WikipediaResponse struct {
	// BatchComplete string `json:"batchcomplete"`
	Query Query `json:"query"`
}

type Query struct {
	SearchInfo SearchInfo `json:"searchinfo"`
	Search     []Search   `json:"search"`
}

type SearchInfo struct {
	Totalhits int `json:"totalhits"`
}

type Search struct {
	Title  string `json:"title"`
	PageId int    `json:"pageid"`
	// Snippet string `json:"snippet"`
}

func main() {
	// 引数チェック
	if len(os.Args) != 2 {
		fmt.Println("検索ワードを指定")
		os.Exit(1)
	}
	arg := os.Args[1]

	// APIを叩くためのURL作成
	baseUrl := url.URL{}
	baseUrl.Scheme = "http"
	baseUrl.Host = fmt.Sprintf("%s.wikipedia.org", "ja")
	baseUrl.Path = "w/api.php"
	query := baseUrl.Query()
	query.Set("action", "query")
	query.Set("list", "search")
	query.Set("srsearch", arg)
	query.Set("srlimit", "10") // 取得する件数
	query.Set("format", "json")
	baseUrl.RawQuery = query.Encode()

	// 記事検索APIを叩く
	resp, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println("Error!")
	}
	defer resp.Body.Close()

	// レスポンスをパースする
	body, err := ioutil.ReadAll(resp.Body)
	wikipediaResponse := new(WikipediaResponse)
	err = json.Unmarshal(body, wikipediaResponse)

	// ヒットした記事が0件の場合は終了
	if wikipediaResponse.Query.SearchInfo.Totalhits <= 0 {
		fmt.Println("記事が見つかりませんでした。")
		os.Exit(1)
	}

	// 記事タイトルとURLを表示
	for _, v := range wikipediaResponse.Query.Search {
		fmt.Println("---------------------------------------------------")
		fmt.Println(v.Title)
		fmt.Printf("https://%s.wikipedia.org/?curid=%d\n", "ja", v.PageId)
	}
	fmt.Println("---------------------------------------------------")

}
