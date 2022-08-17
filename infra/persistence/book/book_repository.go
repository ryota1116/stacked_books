package book

import (
	"github.com/ryota1116/stacked_books/domain/model/book"
	"github.com/ryota1116/stacked_books/handler/http/request/user_book/register_user_books"
	"github.com/ryota1116/stacked_books/infra/persistence"
)

type bookPersistence struct{}

func NewBookPersistence() book.BookRepository {
	return &bookPersistence{}
}

// FindOrCreateByGoogleBooksId : GoogleBooksIDからBookレコードを検索し、存在しなければ作成する
func (bookPersistence) FindOrCreateByGoogleBooksId(requestBody RegisterUserBooks.RequestBody) book.Book {
	db := persistence.DbConnect()
	book := book.Book{}
	db.Where(book.Book{GoogleBooksId: requestBody.Book.GoogleBooksId}).FirstOrCreate(&book)

	return book
}
