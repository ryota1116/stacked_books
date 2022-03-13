package handler

import (
	"github.com/ryota1116/stacked_books/domain/model"
	RegisterUserBooks "github.com/ryota1116/stacked_books/handler/http/request/user_book/register_user_books"
	"github.com/ryota1116/stacked_books/tests/test_assertion"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// テストで期待するレスポンスボディJSON文字列のファイルパス
const expectedRegisterUserBookJson = "../tests/expected/api/userBookHandler/200_register_user_book.json"

type UserBookUseCaseMock struct {}

func (UserBookUseCaseMock) RegisterUserBook(int, RegisterUserBooks.RequestBody) (model.Book, model.UserBook) {
	return model.Book{
			Id:             1,
			GoogleBooksId:  "Wx1dLwEACAAJ",
			Title:          "リーダブルコード",
			Description:    "読んでわかるコードの重要性と方法について解説",
			Image:          "",
			Isbn_10:         "4873115655",
			Isbn_13:         "9784873115658",
			PageCount:      237,
			PublishedYear:  2012,
			PublishedMonth: 6,
			PublishedDate:  0,
			CreatedAt:      time.Time{},
			UpdatedAt:      time.Time{},
	}, model.UserBook{
			Id:        1,
			UserId:    1,
			BookId:    1,
			Status:    1,
			Memo:      "メモメモメモ",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			Book:      model.Book{},
		}
}

func (m UserBookUseCaseMock) FindUserBooksByUserId(userId int) ([]model.Book, error) {
	// 現状テストで使ってないから、空のままにしている
	panic("implement me")
}

type UserSessionHandlerMiddleWareMock struct {}

func (UserSessionHandlerMiddleWareMock) CurrentUser(*http.Request) model.User {
	return model.User{
		Id:        1,
		UserName:  "",
		Email:     "",
		Password:  "",
		Avatar:    "",
		Role:      0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
		Books:     nil,
	}
}

func TestBookHandlerRegisterUserBook(t *testing.T) {
	ubu := UserBookUseCaseMock{}
	ushmw := UserSessionHandlerMiddleWareMock{}
	ubh := NewUserBookHandler(ubu, ushmw)

	// リクエストボディを検証しているならこの記述が活きてくる気がするが、、
	bodyReader := strings.NewReader(`{
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
		"user_book" :{
			"status": 1,
			"memo": "メモメモメモ"
		}
	}`)

	r := httptest.NewRequest("GET", "/register/book", bodyReader)
	w := httptest.NewRecorder()

	r.Header.Add("Authorization", "")

	// この中でUserBookUseCaseMockのRegisterUserBookが実行される
	ubh.RegisterUserBook(w, r)

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

	// JSON文字列の比較
	test_assertion.CompareResponseBodyWithJsonFile(t, responseBodyBytes, expectedRegisterUserBookJson)
}
