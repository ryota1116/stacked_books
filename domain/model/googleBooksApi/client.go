package googleBooksApi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// URLForGoogleBooksAPI :
const URLForGoogleBooksAPI = "https://www.googleapis.com/books/v1/volumes?q="

type Client struct {

}

// SendRequest : GoogleBooksAPIにリクエストを送信する
func (client Client) SendRequest(searchWord string) (SearchBooksResponses, error) {
	// TODO: 以下の方法で「文字列を連結してURLを生成」したいけど上手くいかない
	//var byteURL = make([]byte, 0, 100) // 100byte のキャパシティを確保
	//byteURL = append(byteURL, []byte(URLForGoogleBooksAPI))
	//byteURL = append(byteURL, searchWord[0])
	//searchURL := string(byteURL)

	// 文字列を連結してURLを生成
	searchURL := URLForGoogleBooksAPI + searchWord

	// GoogleBooksAPIを叩く
	res, err := http.Get(searchURL)

	if err != nil {
		fmt.Println(err)
		return SearchBooksResponses{}, err
		// fmt.Errorf("Unable to get this url : http status %d", res.StatusCode)
	}

	// レスポンスボディをcloseする
	// deferは関数終了時(return時)に実行される
	// https://blog.web-apps.tech/net-http-for-valid/
	// https://qiita.com/stk0724/items/dc400dccd29a4b3d6471
	defer res.Body.Close()

	// レスポンスボディを読み込む
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return SearchBooksResponses{}, err
	}

	// JSONエンコードされたデータをparseして、構造体に格納する
	var responseFromGoogleBooksAPI ResponseBodyFromGoogleBooksAPI
	if err := json.Unmarshal(body, &responseFromGoogleBooksAPI); err != nil {
		return SearchBooksResponses{}, err
	}

	// GoogleBooksAPIのJSONレスポンスの構造体から、 書籍検索用のレスポンスボディ構造体を生成する
	searchBooksResponse := SearchBooksResponseGenerator{
		ResponseBodyFromGoogleBooksAPI: responseFromGoogleBooksAPI,
	}.execute()

	// レスポンスボディを返す
	return searchBooksResponse, nil
}
