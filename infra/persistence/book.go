package persistence

import (
	"github.com/ryota1116/stacked_books/domain/model"
	"github.com/ryota1116/stacked_books/domain/repository"
	"github.com/ryota1116/stacked_books/handler/http/request/user_book/register_user_books"
)

type bookPersistence struct {}

func NewBookPersistence() repository.BookRepository {
	return &bookPersistence{}
}

// FindOrCreateByGoogleBooksId : GoogleBooksIDからBookレコードを検索し、存在しなければ作成する
func (bookPersistence) FindOrCreateByGoogleBooksId(requestBody RegisterUserBooks.RequestBody) model.Book {
	db := DbConnect()
	book := model.Book{}
	db.Where(model.Book{GoogleBooksId: requestBody.Book.GoogleBooksId}).FirstOrCreate(&book)

	return book
}
