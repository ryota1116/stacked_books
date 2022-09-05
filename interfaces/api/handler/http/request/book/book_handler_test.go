package book

import (
	"encoding/json"
	"github.com/magiconair/properties/assert"
	"github.com/ryota1116/stacked_books/domain/model/google-books-api"
	res "github.com/ryota1116/stacked_books/interfaces/api/handler/http/response"
	"github.com/ryota1116/stacked_books/tests/test_assertion"
	"io/ioutil"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

// テストで期待するレスポンスボディJSON文字列のファイルパス
const expectedSearchBooksJson = "../tests/expected/api/book_handler/200_search_books_response.json"

// bookUseCaseMock : BookUseCaseInterfaceを実装しているモック
type bookUseCaseMock struct{}

// SearchBooks : インターフェイスを満たすためのメソッド
func (bu bookUseCaseMock) SearchBooks(title string) (google_books_api.ResponseBodyFromGoogleBooksAPI, error) {
	return google_books_api.ResponseBodyFromGoogleBooksAPI{
		Items: []google_books_api.Item{
			{
				ID: "Wx1dLwEACAAJ",
				VolumeInfo: google_books_api.VolumeInfo{
					Title:         "リーダブルコード",
					Authors:       []string{"Dustin Boswell", "Trevor Foucher"},
					PublishedDate: "2012-06",
					Description:   "読んでわかるコードの重要性と方法について解説",
					IndustryIdentifiers: []google_books_api.IndustryIdentifier{
						{
							Type:       "ISBN_10",
							Identifier: "4873115655",
						},
						{
							Type:       "ISBN_13",
							Identifier: "9784873115658",
						},
					},
					PageCount: 237,
				},
			},
			{
				ID: "n6YqDwAAQBAJ",
				VolumeInfo: google_books_api.VolumeInfo{
					Title:         "ExcelVBAを実務で使い倒す技術",
					Authors:       []string{"高橋宣成"},
					PublishedDate: "2017-04",
					Description:   "本書では、VBAを実務の現場で活かすための知識(テクニック)と知恵(考え方とコツ)を教えます!",
					IndustryIdentifiers: []google_books_api.IndustryIdentifier{
						{
							Type:       "ISBN_10",
							Identifier: "4798049999",
						},
						{
							Type:       "ISBN_13",
							Identifier: "9784798049991",
						},
					},
					PageCount: 289,
				},
			},
		},
	}, nil
}

func TestMain(m *testing.M) {
	// テストコードの実行（testing.M.Runで各テストケースが実行され、成功の場合0を返す）
	// また、各ユニットテストの中でテストデータをinsertすれば良さそう。
	status := m.Run()

	// 0が渡れば成功する。プロセスのkillも実行される。
	os.Exit(status)
}

// 外部APIを用いた書籍検索のエンドポイントのテスト
func TestBookHandler_SearchBooks(t *testing.T) {
	bh := NewBookHandler(bookUseCaseMock{})

	t.Run("正常系のテスト", func(t *testing.T) {
		bodyReader := strings.NewReader(`{
			"title": "リーダブルコード"
		}`)

		r := httptest.NewRequest(
			"GET",
			"/books/search",
			bodyReader)
		w := httptest.NewRecorder()

		// handler/userbook.goのSearchBooksメソッドを呼び出し、
		// その中でbookUseCaseMockのSearchBooksメソッドが呼び出されている
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

		// レスポンスボディのjson文字列をテスト
		test_assertion.CompareResponseBodyWithJsonFile(t, responseBodyBytes, expectedSearchBooksJson)
	})

	t.Run("異常系_リクエストボディにTitleが含まれていない場合", func(t *testing.T) {
		// リクエストボディにTitleが含まれていない場合
		bodyReader := strings.NewReader(`{}`)

		r := httptest.NewRequest("GET", "/books/search", bodyReader)
		w := httptest.NewRecorder()

		bh.SearchBooks(w, r)

		// レスポンスを代入
		response := w.Result()

		// ステータスコードのテスト(バリデーションエラーによりステータスコードが422を期待)
		if response.StatusCode != 422 {
			t.Errorf(`レスポンスのステータスコードは %d でした`, response.StatusCode)
		}

		// レスポンスボディを[]byte型に変換
		responseBodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		// []byte型を構造体に格納
		var errorResponseBody res.ErrorResponseBody
		if err := json.Unmarshal(responseBodyBytes, &errorResponseBody); err != nil {
			panic(err)
		}

		// レスポンスボディの結果をテスト(構造体に戻してテストしている)
		assert.Equal(t, errorResponseBody.Message, "本のタイトルを入力してください")
	})

	t.Run("異常系_リクエストボディのTitleの値が空の場合", func(t *testing.T) {
		// リクエストボディのTitleが空の場合
		bodyReader := strings.NewReader(`{
			"title": ""
		}`)

		r := httptest.NewRequest("GET", "/books/search", bodyReader)
		w := httptest.NewRecorder()

		bh.SearchBooks(w, r)

		// レスポンスを代入
		response := w.Result()

		// NOTE: ステータスコード422はHandlerが返しているから、Handlerの責務としてテストして良さそう。
		// ステータスコードのテスト(バリデーションエラーによりステータスコードが422を期待)
		if response.StatusCode != 422 {
			t.Errorf(`レスポンスのステータスコードは %d でした`, response.StatusCode)
			t.Errorf(`レスポンスボディは「 %s 」でした`, response.Body)
		}

		// レスポンスボディを[]byte型に変換
		responseBodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		// []byte型を構造体に格納
		var errorResponseBody res.ErrorResponseBody
		if err := json.Unmarshal(responseBodyBytes, &errorResponseBody); err != nil {
			panic(err)
		}

		// TODO: バリデーションメッセージはValidatorが返しているから、Handlerの責務ではない。
		// レスポンスボディの結果をテスト(構造体に戻してテストしている)
		assert.Equal(t, errorResponseBody.Message, "本のタイトルを入力してください")
	})
}