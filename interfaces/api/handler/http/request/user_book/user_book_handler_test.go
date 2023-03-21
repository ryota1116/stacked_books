package user_book

import (
	"github.com/ryota1116/stacked_books/tests"
	"github.com/ryota1116/stacked_books/tests/test_assertion"
	bookUseCase "github.com/ryota1116/stacked_books/usecase/book"
	user2 "github.com/ryota1116/stacked_books/usecase/user"
	userBookUseCase "github.com/ryota1116/stacked_books/usecase/userbook"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"
)

// テストで期待するレスポンスボディJSON文字列のファイルパス
const expectedJsonDirectory = "/tests/expected/api/userBookHandler"

type UserBookUseCaseMock struct{}

func (UserBookUseCaseMock) RegisterUserBook(_ userBookUseCase.UserBookCreateCommand) (bookUseCase.BookDto, userBookUseCase.UserBookDto, error) {
	description := "読んでわかるコードの重要性と方法について解説"
	image := ""
	isbn10 := "4873115655"
	isbn13 := "9784873115658"
	publishedYear := 2012
	publishedMonth := 6
	publishedDate := 0
	memo := "メモメモメモ"

	return bookUseCase.BookDto{
			Id:             1,
			GoogleBooksId:  "Wx1dLwEACAAJ",
			Title:          "リーダブルコード",
			Description:    &description,
			Image:          &image,
			Isbn10:         &isbn10,
			Isbn13:         &isbn13,
			PageCount:      237,
			PublishedYear:  &publishedYear,
			PublishedMonth: &publishedMonth,
			PublishedDate:  &publishedDate,
		}, userBookUseCase.UserBookDto{
			Id:        1,
			UserId:    1,
			BookId:    1,
			Status:    1,
			Memo:      &memo,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}, nil
}

func (UserBookUseCaseMock) FindUserBooksByUserId(_ int) ([]bookUseCase.BookDto, error) {
	description := "読んでわかるコードの重要性と方法について解説"
	image := ""
	isbn10 := "4873115655"
	isbn13 := "9784873115658"
	publishedYear := 2012
	publishedMonth := 6
	publishedDate := 0

	var userBookDto []bookUseCase.BookDto
	userBookDto = append(userBookDto, bookUseCase.BookDto{
		Id:             1,
		GoogleBooksId:  "Wx1dLwEACAAJ",
		Title:          "リーダブルコード",
		Description:    &description,
		Image:          &image,
		Isbn10:         &isbn10,
		Isbn13:         &isbn13,
		PageCount:      237,
		PublishedYear:  &publishedYear,
		PublishedMonth: &publishedMonth,
		PublishedDate:  &publishedDate,
	})

	return userBookDto, nil
}

type UserSessionHandlerMiddleWareMock struct{}

func (UserSessionHandlerMiddleWareMock) CurrentUser(_ *http.Request) (user2.UserDto, error) {
	return user2.UserDto{
		Id:       1,
		UserName: "user_name",
		Email:    "user@example.com",
		Password: "password",
	}, nil
}

func TestMain(m *testing.M) {
	// テストコードの実行（testing.M.Runで各テストケースが実行され、成功の場合0を返す）
	// => また各ユニットテストの中でテストデータをinsertすれば良さそう。
	status := m.Run()

	// 0が渡れば成功する。プロセスのkillも実行される。
	os.Exit(status)
}

func TestUserBookHandler_RegisterUserBook(t *testing.T) {
	ubh := NewUserBookHandler(UserBookUseCaseMock{}, UserSessionHandlerMiddleWareMock{})

	// jsonファイルの絶対パスを取得(TODO: ローカル用の取得になっているので修正する)
	_, testFilePath, _, _ := runtime.Caller(0)
	projectRootDir := filepath.Join(filepath.Dir(testFilePath), "..", "..", "..", "..", "..", "..")

	t.Run("正常系のテスト", func(t *testing.T) {
		// リクエスト
		body := strings.NewReader(`{
		"book" :{
			"google_books_id": "Wx1dLwEACAAJ",
			"title": "リーダブルコード",
			"authors": ["Dustin Boswell","Trevor Foucher"],
			"description": "読んでわかるコードの重要性と方法について解説",
			"isbn_10": "4873115655",
			"isbn_13": "9784873115658",
			"page_count": 237,
			"published_year": 2012,
			"published_month": 6
		},
		"userbook" :{
			"status": 1,
			"memo": "メモメモメモ"
			}
		}`)
		r := httptest.NewRequest("GET", "/register/userbook", body)
		w := httptest.NewRecorder()
		r.Header.Add("Authorization", "")
		ubh.RegisterUserBook(w, r)
		response := w.Result()

		// ステータスコードのテスト
		if response.StatusCode != 200 {
			testHandler := tests.TestHandler{T: t}
			testHandler.PrintErrorFormatFromResponse(response)
		}

		expectedJsonFilePath := filepath.Join(
			projectRootDir,
			expectedJsonDirectory+"/register_user_book/200_response.json",
		)

		// レスポンスボディのjson文字列をテスト
		test_assertion.CompareResponseBodyWithJsonFile(
			t,
			response.Body,
			expectedJsonFilePath,
		)
	})
}

func TestUserBookHandler_FindUserBooks(t *testing.T) {
	ubh := NewUserBookHandler(UserBookUseCaseMock{}, UserSessionHandlerMiddleWareMock{})

	// jsonファイルの絶対パスを取得(TODO: ローカル用の取得になっているので修正する)
	_, testFilePath, _, _ := runtime.Caller(0)
	projectRootDir := filepath.Join(filepath.Dir(testFilePath), "..", "..", "..", "..", "..", "..")

	t.Run("正常系のテスト", func(t *testing.T) {
		// リクエスト
		body := strings.NewReader(``)
		r := httptest.NewRequest("GET", "/register/userbook", body)
		w := httptest.NewRecorder()
		r.Header.Add("Authorization", "")
		ubh.FindUserBooks(w, r)
		response := w.Result()

		// ステータスコードのテスト
		if response.StatusCode != 200 {
			testHandler := tests.TestHandler{T: t}
			testHandler.PrintErrorFormatFromResponse(response)
		}

		expectedJsonFilePath := filepath.Join(
			projectRootDir,
			expectedJsonDirectory+"/find_user_books/200_response.json",
		)

		// レスポンスボディのjson文字列をテスト
		test_assertion.CompareResponseBodyWithJsonFile(
			t,
			response.Body,
			expectedJsonFilePath,
		)
	})
}
