package book

import (
	"fmt"
	"github.com/ryota1116/stacked_books/domain/model/searched_books/google_books_api"
	"github.com/ryota1116/stacked_books/tests"
	"github.com/ryota1116/stacked_books/usecase/book"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// テストで期待するレスポンスボディJSON文字列のファイルパス
const expectedJsonDirectory = "/tests/expected/api/book_handler"

// BookUseCaseInterfaceを実装しているモック構造体
type bookUseCaseMock struct{}

// インターフェイスを満たすためのメソッド
func (bu bookUseCaseMock) SearchBooks(string) (google_books_api.ResponseBodyFromGoogleBooksApi, error) {
	return google_books_api.ResponseBodyFromGoogleBooksApi{
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

func (bu bookUseCaseMock) GetBookById(_ int) (book.Dto, error) {
	description := "読んでわかるコードの重要性と方法について解説"
	image := ""
	isbn10 := "4873115655"
	isbn13 := "9784873115658"
	publishedYear := 2012
	publishedMonth := 6
	publishedDate := 0

	return book.Dto{
		Id:             1,
		GoogleBooksId:  "test_id",
		Title:          "タイトル",
		Description:    &description,
		Image:          &image,
		Isbn10:         &isbn10,
		Isbn13:         &isbn13,
		PageCount:      100,
		PublishedYear:  &publishedYear,
		PublishedMonth: &publishedMonth,
		PublishedDate:  &publishedDate,
	}, nil
}

func TestMain(m *testing.M) {
	// テストコードの実行（testing.M.Runで各テストケースが実行され、成功の場合0を返す）
	// => また各ユニットテストの中でテストデータをinsertすれば良さそう。
	status := m.Run()

	// 0が渡れば成功する。プロセスのkillも実行される。
	os.Exit(status)
}

// 外部APIを用いた書籍検索のエンドポイントのテスト
func TestBookHandler_SearchBooks(t *testing.T) {
	// モックを注入している
	bh := NewBookHandler(bookUseCaseMock{})

	// jsonファイルの絶対パスを取得(TODO: ローカル用の取得になっているので修正する)
	_, testFilePath, _, _ := runtime.Caller(0)
	projectRootDir := filepath.Join(filepath.Dir(testFilePath), "..", "..", "..", "..", "..", "..")

	t.Run("正常系のテスト", func(t *testing.T) {
		testHandler := tests.TestHandler{T: t}

		title := "リーダブルコード"
		body := strings.NewReader(``)

		// リクエスト
		r := httptest.NewRequest(
			"GET",
			fmt.Sprintf("/books/search?title=%s", title),
			body,
		)
		w := httptest.NewRecorder()
		bh.SearchBooks(w, r)   // この中でbookUseCaseMockのSearchBooksメソッドが呼び出される
		response := w.Result() // レスポンスを代入

		// ステータスコードのテスト
		if response.StatusCode != 200 {
			testHandler.PrintErrorFormatFromResponse(response)
		}

		expectedJsonFilePath := filepath.Join(
			projectRootDir,
			expectedJsonDirectory+"/search_books/200_response.json",
		)

		// レスポンスボディのjson文字列をテスト
		testHandler.CompareResponseBodyWithJsonFile(
			response.Body,
			expectedJsonFilePath,
		)
	})

	t.Run("異常系_リクエストボディにTitleが含まれていない場合", func(t *testing.T) {
		testHandler := tests.TestHandler{T: t}

		bodyReader := strings.NewReader(`{}`)
		r := httptest.NewRequest("GET", "/books/search", bodyReader)
		w := httptest.NewRecorder()
		bh.SearchBooks(w, r)
		response := w.Result()

		// ステータスコードのテスト(バリデーションエラーによりステータスコードが422を期待)
		if response.StatusCode != 422 {
			testHandler.PrintErrorFormatFromResponse(response)
		}

		expectedJsonFilePath := filepath.Join(
			projectRootDir,
			expectedJsonDirectory+"/search_books/response_without_title.json",
		)

		// レスポンスボディのjson文字列をテスト
		testHandler.CompareResponseBodyWithJsonFile(
			response.Body,
			expectedJsonFilePath,
		)
	})

	t.Run("異常系_リクエストボディのTitleの値が空の場合", func(t *testing.T) {
		testHandler := tests.TestHandler{T: t}

		bodyReader := strings.NewReader(`{
			"title": ""
		}`)
		r := httptest.NewRequest("GET", "/books/search", bodyReader)
		w := httptest.NewRecorder()
		bh.SearchBooks(w, r)
		response := w.Result()

		// NOTE: ステータスコード422はHandlerが返しているから、Handlerの責務としてテストして良さそう。
		// ステータスコードのテスト(バリデーションエラーによりステータスコードが422を期待)
		if response.StatusCode != 422 {
			testHandler.PrintErrorFormatFromResponse(response)
		}

		expectedJsonFilePath := filepath.Join(
			projectRootDir,
			expectedJsonDirectory+"/search_books/response_without_title.json",
		)

		// レスポンスボディのjson文字列をテスト
		testHandler.CompareResponseBodyWithJsonFile(
			response.Body,
			expectedJsonFilePath,
		)
	})
}
