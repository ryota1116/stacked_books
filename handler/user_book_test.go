package handler

import (
	"github.com/ryota1116/stacked_books/domain/model"
	"testing"
	"time"
)

type UserBookUseCaseMock struct {}

func (uu *UserUseCaseMock) RegisterUserBook(userBookParameter model.UserBookParameter) model.UserBookParameter {
	return model.UserBookParameter{
		UserBook: model.UserBook{
			Id:        1,
			UserId:    1,
			BookId:    1,
			Status:    0,
			Memo:      "",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Book: model.Book{
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
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		},
	}
}

func (uu *UserUseCaseMock) ReadUserBooks(userId int) []model.Book {
	// スライスの作成
	books := make([]model.Book, 1)

	book := model.Book{
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
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	return append(books, book)
}

func (uu *UserUseCaseMock) GetUserTotalReadingVolume(userId int) int {
	return 100
}

func TestUserHandlerRegisterUserBookWithoutBookTitle(t *testing.T) {

}

func TestUserHandlerRegisterUserBookWithInvalidIsbn10(t *testing.T) {

}

func TestUserHandlerReadUserBooks(t *testing.T) {

}

func TestUserHandlerGetUserTotalReadingVolume(t *testing.T) {

}
