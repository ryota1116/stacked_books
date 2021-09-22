package handler

import (
	"github.com/ryota1116/stacked_books/domain/model"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type UserBookUseCaseMock struct {}

func (ubum UserBookUseCaseMock) RegisterUserBook(userId int, userBookParameter model.UserBookParameter) model.UserBookParameter {
	return model.UserBookParameter{
		UserBook: model.UserBook{
			Id:        1,
			UserId:    1,
			BookId:    1,
			Status:    0,
			Memo:      "メモメモメモ",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
		Book:     model.Book{
			Id:             1,
			GoogleBooksId:  "Wx1dLwEACAAJ",
			Title:          "リーダブルコード",
			Description:    "読んでわかるコードの重要性と方法について解説",
			Isbn_10:        "4873115655",
			Isbn_13:        "9784873115658",
			PageCount:      237,
			PublishedYear:  2012,
			PublishedMonth: 6,
			CreatedAt:      time.Time{},
			UpdatedAt:      time.Time{},
		},
	}
}

func TestBookHandlerRegisterUserBook(t *testing.T) {
	ubu := UserBookUseCaseMock{}
	ubh := NewUserBookHandler(ubu)

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
		"memo": "めもめもめも"
	}`)

	r := httptest.NewRequest(
		"GET",
		"/register/book",
		bodyReader)
	w := httptest.NewRecorder()

	ubh.RegisterUserBook(w, r)
}
