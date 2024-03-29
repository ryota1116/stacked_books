package book

import (
	"encoding/json"
	"github.com/ryota1116/stacked_books/domain/model/searched_books/google_books_api"
	"github.com/ryota1116/stacked_books/tests"
	"testing"
)

// GoogleBooksAPIのJSONレスポンスの構造体から、書籍検索用のレスポンスボディ構造体を生成するメソッドの正常系テスト
func TestSearchBooksResponseGeneratorExecute(t *testing.T) {
	// 書籍検索用レスポンスボディ構造体の期待値を作る
	var expectedSearchBooksResponse []SearchBooksResponse
	expectedSearchBooksResponse = append(expectedSearchBooksResponse, SearchBooksResponse{
		GoogleBooksId: "Wx1dLwEACAAJ",
		Title:         "リーダブルコード",
		Authors:       []string{"Dustin Boswell", "Trevor Foucher"},
		Description:   "読んでわかるコードの重要性と方法について解説",
		Isbn10:        "4873115655",
		Isbn13:        "9784873115658",
		PageCount:     237,
		RegisteredAt:  "2012-06",
	})
	expected := SearchBooksResponses{expectedSearchBooksResponse}

	// GoogleBooksAPIが返すJSONレスポンス
	responseBody := []byte(`
	{
	  "kind": "books#volumes",
	  "totalItems": 1647,
	  "items": [
		{
		  "kind": "books#volume",
		  "id": "Wx1dLwEACAAJ",
		  "etag": "2i+rhAaWyC4",
		  "selfLink": "https://www.googleapis.com/books/v1/volumes/Wx1dLwEACAAJ",
		  "volumeInfo": {
			"title": "リーダブルコード",
			"subtitle": "より良いコードを書くためのシンプルで実践的なテクニック",
			"authors": [
			  "Dustin Boswell",
			  "Trevor Foucher"
			],
			"publisher": "O'Reilly Media, Inc.",
			"publishedDate": "2012-06",
			"description": "読んでわかるコードの重要性と方法について解説",
			"industryIdentifiers": [
			  {
				"type": "ISBN_10",
				"identifier": "4873115655"
			  },
			  {
				"type": "ISBN_13",
				"identifier": "9784873115658"
			  }
			],
			"readingModes": {
			  "text": false,
			  "image": false
			},
			"pageCount": 237,
			"printType": "BOOK",
			"averageRating": 5,
			"ratingsCount": 1,
			"maturityRating": "NOT_MATURE",
			"allowAnonLogging": false,
			"contentVersion": "preview-1.0.0",
			"imageLinks": {
			  "smallThumbnail": "http://books.google.com/books/content?id=Wx1dLwEACAAJ&printsec=frontcover&img=1&zoom=5&source=gbs_api",
			  "thumbnail": "http://books.google.com/books/content?id=Wx1dLwEACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api"
			},
			"language": "ja",
			"previewLink": "http://books.google.co.jp/books?id=Wx1dLwEACAAJ&dq=%E3%83%AA%E3%83%BC%E3%83%80%E3%83%96%E3%83%AB%E3%82%B3%E3%83%BC%E3%83%89&hl=&cd=1&source=gbs_api",
			"infoLink": "http://books.google.co.jp/books?id=Wx1dLwEACAAJ&dq=%E3%83%AA%E3%83%BC%E3%83%80%E3%83%96%E3%83%AB%E3%82%B3%E3%83%BC%E3%83%89&hl=&source=gbs_api",
			"canonicalVolumeLink": "https://books.google.com/books/about/%E3%83%AA%E3%83%BC%E3%83%80%E3%83%96%E3%83%AB%E3%82%B3%E3%83%BC%E3%83%89.html?hl=&id=Wx1dLwEACAAJ"
		  },
		  "saleInfo": {
			"country": "JP",
			"saleability": "NOT_FOR_SALE",
			"isEbook": false
		  },
		  "accessInfo": {
			"country": "JP",
			"viewability": "NO_PAGES",
			"embeddable": false,
			"publicDomain": false,
			"textToSpeechPermission": "ALLOWED",
			"epub": {
			  "isAvailable": false
			},
			"pdf": {
			  "isAvailable": false
			},
			"webReaderLink": "http://play.google.com/books/reader?id=Wx1dLwEACAAJ&hl=&printsec=frontcover&source=gbs_api",
			"accessViewStatus": "NONE",
			"quoteSharingAllowed": false
		  },
		  "searchInfo": {
			"textSnippet": "読んでわかるコードの重要性と方法について解説"
		  }
		}
      ]
	}`)

	// JSONデータをparseして、構造体に格納する
	var responseFromGoogleBooksAPI google_books_api.ResponseBodyFromGoogleBooksApi
	if err := json.Unmarshal(responseBody, &responseFromGoogleBooksAPI); err != nil {
		t.Errorf(err.Error())
	}

	// 書籍検索用のレスポンスボディ構造体を生成する(テスト対象)
	searchBooksResponses := SearchBooksResponseGenerator{
		ResponseBodyFromGoogleBooksApi: responseFromGoogleBooksAPI,
	}.Execute()

	// 戻り値である構造体が正しいことをテスト
	tests.Assertion{T: t}.AssertEqual(expected, searchBooksResponses)

	// TODO: 差分がある場合に、cmpを使って差分を出力できるようにしたい
	//if diff := cmp.Diff(book, expectedBook); diff != "" {
	//	t.Errorf("戻り値の構造体が期待するものではありません。: (-got +want)\n%s", diff)
	//}
}
