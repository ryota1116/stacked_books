package userbook

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/ryota1116/stacked_books/domain/model/book"
	"github.com/ryota1116/stacked_books/domain/model/userbook"
	RegisterUserBooks "github.com/ryota1116/stacked_books/handler/http/request/user_book/register_user_books"
	"github.com/ryota1116/stacked_books/tests/expected/api/user_book_use_case"
	"strings"
	"testing"
	"time"
)

type BookRepositoryMock struct{}

func (BookRepositoryMock) FindOrCreateByGoogleBooksId(body RegisterUserBooks.RequestBody) book.Book {
	return book.Book{
		GoogleBooksId:  "Wx1dLwEACAAJ",
		Title:          "リーダブルコード",
		Description:    "読んでわかるコードの重要性と方法について解説",
		Isbn_10:        "4873115655",
		Isbn_13:        "9784873115658",
		PageCount:      237,
		PublishedYear:  2012,
		PublishedMonth: 6,
		PublishedDate:  0,
	}
}

type UserBookRepositoryMock struct{}

func (UserBookRepositoryMock) CreateOne(userId int, bookId int, requestBody RegisterUserBooks.RequestBody) userbook.UserBook {
	return userbook.UserBook{
		Id:     1,
		UserId: 1,
		BookId: 1,
		Status: 1,
		Memo:   "メモメモメモ",
	}
}

func (UserBookRepositoryMock) FindAllByUserId(userId int) ([]book.Book, error) {
	var books []book.Book
	books = append(books, book.Book{
		Id:             1,
		GoogleBooksId:  "test",
		Title:          "タイトル",
		Description:    "説明文です",
		Image:          "",
		Isbn_10:        "",
		Isbn_13:        "",
		PageCount:      100,
		PublishedYear:  2022,
		PublishedMonth: 8,
		PublishedDate:  10,
		CreatedAt:      time.Date(2022, time.August, 10, 12, 0, 0, 0, time.UTC),
		UpdatedAt:      time.Date(2022, time.August, 10, 12, 0, 0, 0, time.UTC),
		Users:          nil,
	})

	return books, nil
}

// UserBookUseCaseのRegisterUserBookの正常系テスト
func TestUserBookUseCaseRegisterUserBook(t *testing.T) {
	brm := BookRepositoryMock{}
	ubrm := UserBookRepositoryMock{}
	ubu := NewUserBookUseCase(brm, ubrm)

	bodyReader := strings.NewReader(`{
		"google_books_id": "Wx1dLwEACAAJ",
		"title": "リーダブルコード",
		"authors": ["Dustin Boswell","Trevor Foucher"],
		"description": "読んでわかるコードの重要性と方法について解説",
		"isbn_10": "4873115655",
		"isbn_13": "9784873115658",
		"page_count": 237,
		"published_year": 2012,
		"published_month": 6,
	
		"status": 0,
		"memo": "メモメモメモ"
	}`)

	// json文字列をリクエストボディに変換
	requestBody := RegisterUserBooks.RequestBody{}
	err := json.NewDecoder(bodyReader).Decode(&requestBody)
	if err != nil {
		fmt.Println(err)
	}

	// userBookUseCaseのRegisterUserBookを実行
	book, userBook := ubu.RegisterUserBook(1, requestBody)

	// 戻り値である構造体が正しいことをテスト
	if diff := cmp.Diff(book, user_book_use_case.ExpectedBookStructForRegisterUserBook); diff != "" {
		t.Errorf("戻り値の構造体が期待するものではありません。: (-got +want)\n%s", diff)
	}

	if diff := cmp.Diff(userBook, user_book_use_case.ExpectedUserBookStructForRegisterUserBook); diff != "" {
		t.Errorf("戻り値の構造体が期待するものではありません。: (-got +want)\n%s", diff)
	}
}
