package persistence

import (
	"github.com/ryota1116/stacked_books/domain/model"
	"github.com/ryota1116/stacked_books/domain/model/dto"
	"github.com/ryota1116/stacked_books/domain/repository"
)

type bookPersistence struct {}

func NewBookPersistence() repository.BookRepository {
	return &bookPersistence{}
}

// FindOrCreateByGoogleBooksId : GoogleBooksIDからBookレコードを検索し、存在しなければ作成する
func (bookPersistence) FindOrCreateByGoogleBooksId(registerUserBookRequestParameter dto.RegisterUserBookRequestParameter) model.Book {
	db := DbConnect()
	book := model.Book{}
	db.Where(model.Book{GoogleBooksId: registerUserBookRequestParameter.Book.GoogleBooksId}).FirstOrCreate(&book)

	return book
}
