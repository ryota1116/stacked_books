package repository

import (
	"github.com/ryota1116/stacked_books/domain/model"
	"github.com/ryota1116/stacked_books/handler/http/request/user_book/register_user_books"
)

type BookRepository interface {
	FindOrCreateByGoogleBooksId(parameter RegisterUserBooks.RequestBody) model.Book
}
