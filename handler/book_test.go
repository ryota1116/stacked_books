package handler

import (
	"encoding/json"
	"github.com/ryota1116/stacked_books/domain/model/googleBooksApi"
	"github.com/ryota1116/stacked_books/tests/test_assertion"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

// テストで期待するレスポンスボディJSON文字列のファイルパス
const expectedSearchBooksJson = "../tests/expected/api/userBookHandler/200_search_books.json"

// BookUseCaseMock : BookUseCaseInterfaceを実装しているモック
type BookUseCaseMock struct {}

func (bu BookUseCaseMock) SearchBooks(requestParameter googleBooksApi.RequestParameter) (googleBooksApi.SearchBooksResponses, error) {
	return googleBooksApi.SearchBooksResponses{
		{
			Title:        "リーダブルコード",
			Authors:      []string{"Dustin Boswell", "Trevor Foucher"},
			Description:  "読んでわかるコードの重要性と方法について解説",
			Isbn10:       "4873115655",
			Isbn13:       "9784873115658",
			PageCount:    237,
			RegisteredAt: "2012-06",
		},
		{
			Title:        "ExcelVBAを実務で使い倒す技術",
			Authors:      []string{"高橋宣成"},
			Description:  "本書では、VBAを実務の現場で活かすための知識(テクニック)と知恵(考え方とコツ)を教えます!",
			Isbn10:       "4798049999",
			Isbn13:       "9784798049991",
			PageCount:    289,
			RegisteredAt: "2017-04",
		},
	}, nil
}

// TestBookHandlerSearchBooks : Handler層のSearchBooksメソッドの正常系テスト
func TestBookHandlerSearchBooks(t *testing.T) {
	bu := BookUseCaseMock{}
	bh := NewBookHandler(bu)

	bodyReader := strings.NewReader(`{
		"title": "リーダブルコード"
	}`)

	r := httptest.NewRequest("GET", "/books/search", bodyReader)
	w := httptest.NewRecorder()

	// handler/book.goのSearchBooksメソッドを呼び出し、
	// その中でBookUseCaseMockのSearchBooksメソッドが呼び出されている
	bh.SearchBooks(w, r)

	// レスポンスを代入
	response := w.Result()

	// ステータスコードのテスト
	if response.StatusCode != 200 {
		t.Errorf(`レスポンスのステータスコードは %d でした`, response.StatusCode)
	}

	// レスポンスボディを[]byte型に変換
	responseBodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	// []byte型を構造体に格納
	var searchBooksResponses googleBooksApi.SearchBooksResponses
	if err := json.Unmarshal(responseBodyBytes, &searchBooksResponses); err != nil {
		panic(err)
	}

	// レスポンスボディの結果をテスト
	test_assertion.CompareResponseBodyWithJsonFile(t, responseBodyBytes, expectedSearchBooksJson)
}

// TestBookHandlerSearchBooksWithoutRequestBody : リクエストボディにTitleが含まれていない場合
func TestBookHandlerSearchBooksWithoutRequestBody(t *testing.T) {
	bu := BookUseCaseMock{}
	bh := NewBookHandler(bu)

	// リクエストボディにTitleが含まれていない場合
	bodyReader := strings.NewReader(`{}`)

	r := httptest.NewRequest("GET", "/books/search", bodyReader)
	w := httptest.NewRecorder()

	bh.SearchBooks(w, r)

	// レスポンスを代入
	response := w.Result()

	// ステータスコードのテスト
	if response.StatusCode != 200 {
		t.Errorf(`レスポンスのステータスコードは %d でした`, response.StatusCode)
	}

	// レスポンスボディを[]byte型に変換
	responseBodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	// []byte型を構造体に格納
	var searchBooksResponses googleBooksApi.SearchBooksResponses
	if err := json.Unmarshal(responseBodyBytes, &searchBooksResponses); err != nil {
		panic(err)
	}

	// レスポンスボディの結果をテスト
	test_assertion.CompareResponseBodyWithJsonFile(t, responseBodyBytes, expectedSearchBooksJson)
}

// TestBookHandlerSearchBooksWithEmptyTitleParameter : リクエストボディのTitleの値が空の場合
func TestBookHandlerSearchBooksWithEmptyParameter(t *testing.T) {

}
