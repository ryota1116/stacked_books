package user_book

import (
	"github.com/ryota1116/stacked_books/domain/model/book"
	"github.com/ryota1116/stacked_books/domain/model/user"
	"github.com/ryota1116/stacked_books/domain/model/userbook"
	"github.com/ryota1116/stacked_books/tests/test_assertion"
	userBookUseCase "github.com/ryota1116/stacked_books/usecase/userbook"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// テストで期待するレスポンスボディJSON文字列のファイルパス
const expectedRegisterUserBookJson = "../tests/expected/api/userBookHandler/200_register_user_book.json"

type UserBookUseCaseMock struct{}

// モック型でプロダクションコードの
func (UserBookUseCaseMock) RegisterUserBook(command userBookUseCase.UserBookCreateCommand) (book.Book, userbook.UserBook) {
	return book.Book{
			Id:             1,
			GoogleBooksId:  "Wx1dLwEACAAJ",
			Title:          "リーダブルコード",
			Description:    "読んでわかるコードの重要性と方法について解説",
			Image:          "",
			Isbn_10:        "4873115655",
			Isbn_13:        "9784873115658",
			PageCount:      237,
			PublishedYear:  2012,
			PublishedMonth: 6,
			PublishedDate:  0,
			CreatedAt:      time.Time{},
			UpdatedAt:      time.Time{},
		}, userbook.UserBook{
			Id:        1,
			UserId:    1,
			BookId:    1,
			Status:    1,
			Memo:      "メモメモメモ",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			Book:      book.Book{},
		}
}

func (UserBookUseCaseMock) FindUserBooksByUserId(userId int) ([]userBookUseCase.UserBookDto, error) {
	var userBookDto []userBookUseCase.UserBookDto
	userBookDto = append(userBookDto, userBookUseCase.UserBookDto{
		ID:             1,
		GoogleBooksId:  "",
		Title:          "",
		Description:    "",
		Isbn10:         "",
		Isbn13:         "",
		PageCount:      0,
		PublishedYear:  0,
		PublishedMonth: 0,
		PublishedDate:  0,
	})

	return userBookDto, nil
}

type UserSessionHandlerMiddleWareMock struct{}

func (UserSessionHandlerMiddleWareMock) CurrentUser(*http.Request) user.User {
	return user.User{
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
		"userbook" :{
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

	r := httptest.NewRequest("GET", "/register/userbook", bodyReader)
	w := httptest.NewRecorder()

	r.Header.Add("Authorization", "")

	// この中でUserBookUseCaseMockのRegisterUserBookが実行される
	ubh.RegisterUserBook(w, r)

	// レスポンスを代入
	response := w.Result()

	// ステータスコードのテスト
	if response.StatusCode != 200 {
		t.Errorf(`レスポンスのステータスコードは %d でした`, response.StatusCode)
		t.Errorf(`レスポンスボディは「 %s 」でした`, response.Body)
	}

	// レスポンスボディを[]byte型に変換
	responseBodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// JSON文字列の比較
	test_assertion.CompareResponseBodyWithJsonFile(t, responseBodyBytes, expectedRegisterUserBookJson)
}
