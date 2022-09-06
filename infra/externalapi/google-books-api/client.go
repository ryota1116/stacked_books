package google_books_api

import (
	"encoding/json"
	model "github.com/ryota1116/stacked_books/domain/model/searched_books/google_books_api"
	"io/ioutil"
	"net/http"
)

const urlForGoogleBooksAPI = "https://www.googleapis.com/books/v1/volumes?q="

type googleBooksApiClient struct{}

func NewClient() model.GoogleBooksApiClientInterface {
	return googleBooksApiClient{}
}

// SendRequest : GoogleBooksAPIにリクエストを送信する
func (client googleBooksApiClient) SendRequest(searchWord string) (model.ResponseBodyFromGoogleBooksApi, error) {
	// TODO: 以下の方法で「文字列を連結してURLを生成」したいけど上手くいかない
	//var byteURL = make([]byte, 0, 100) // 100byte のキャパシティを確保
	//byteURL = append(byteURL, []byte(URLForGoogleBooksAPI))
	//byteURL = append(byteURL, searchWord[0])
	//searchURL := string(byteURL)

	// 文字列を連結してURLを生成
	searchURL := urlForGoogleBooksAPI + searchWord

	// GoogleBooksAPIを叩く
	res, err := http.Get(searchURL)

	if err != nil {
		// return ResponseBodyFromGoogleBooksAPI{}, err
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
		// return model.ResponseBodyFromGoogleBooksAPI{}, err
	}

	// JSONエンコードされたデータをparseして、構造体に格納する
	var responseBodyFromGoogleBooksApi model.ResponseBodyFromGoogleBooksApi
	if err := json.Unmarshal(body, &responseBodyFromGoogleBooksApi); err != nil {
		return model.ResponseBodyFromGoogleBooksApi{}, err
	}

	return responseBodyFromGoogleBooksApi, nil
}
