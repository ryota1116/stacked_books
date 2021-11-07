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
	book := model.Book{
		Title:          registerUserBookRequestParameter.Book.Title,
		Description:    registerUserBookRequestParameter.Book.Description,
		Isbn_10:        registerUserBookRequestParameter.Book.Isbn_10,
		Isbn_13:        registerUserBookRequestParameter.Book.Isbn_13,
		PageCount:      registerUserBookRequestParameter.Book.PageCount,
		PublishedYear:  registerUserBookRequestParameter.Book.PublishedYear,
		PublishedMonth: registerUserBookRequestParameter.Book.PublishedMonth,
		PublishedDate:  registerUserBookRequestParameter.Book.PublishedDate,
	}
	db.Where(model.Book{GoogleBooksId: registerUserBookRequestParameter.Book.GoogleBooksId}).FirstOrCreate(&book)

	return book
}
