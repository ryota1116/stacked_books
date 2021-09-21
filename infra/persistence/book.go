package persistence

import (
	"github.com/ryota1116/stacked_books/domain/model"
	"github.com/ryota1116/stacked_books/domain/repository"
)

type bookPersistence struct {}

func NewBookPersistence() repository.BookRepository {
	return &bookPersistence{}
}

// FindOrCreateByGoogleBooksId : GoogleBooksIDからBookレコードを検索し、存在しなければ作成する
func (bookPersistence) FindOrCreateByGoogleBooksId(googleBooksId string, userBook model.UserBookParameter) model.Book {
	db := DbConnect()
	db.Where("google_books_id = ?", googleBooksId).FirstOrCreate(&userBook.Book)

	return userBook.Book
}

