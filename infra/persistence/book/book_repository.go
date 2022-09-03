package book

import (
	"github.com/ryota1116/stacked_books/domain/model/book"
	"github.com/ryota1116/stacked_books/infra/persistence"
)

type bookPersistence struct{}

func NewBookPersistence() book.BookRepository {
	return &bookPersistence{}
}

// FindOrCreateByGoogleBooksId : GoogleBooksIDからBookレコードを検索し、存在しなければ作成する
func (bookPersistence) FindOrCreateByGoogleBooksId(GoogleBooksId string) book.Book {
	db := persistence.DbConnect()
	b := book.Book{}
	db.Where("google_books_id = ?", GoogleBooksId).FirstOrCreate(&b)

	return b
}
