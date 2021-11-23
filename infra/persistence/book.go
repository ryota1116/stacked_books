package persistence

import (
	"github.com/ryota1116/stacked_books/domain/model"
	"github.com/ryota1116/stacked_books/domain/model/dto"
	"github.com/ryota1116/stacked_books/domain/repository"
	"github.com/ryota1116/stacked_books/infra"
)

type bookPersistence struct {}

func NewBookPersistence() repository.BookRepository {
	return &bookPersistence{}
}

// FindOrCreateByGoogleBooksId : GoogleBooksIDからBookレコードを検索し、存在しなければ作成する
func (bookPersistence) FindOrCreateByGoogleBooksId(registerUserBookRequestParameter dto.RegisterUserBookRequestParameter) model.Book {
	book := model.Book{}
	infra.Db.Where(model.Book{GoogleBooksId: registerUserBookRequestParameter.Book.GoogleBooksId}).FirstOrCreate(&book)

	return book
}
