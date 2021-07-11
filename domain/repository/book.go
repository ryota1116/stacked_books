package repository

import "github.com/ryota1116/stacked_books/domain/model"

type BookRepository interface {
	FindOrCreateByGoogleBooksId(googleBooksId string, userBook model.UserBookParameter) model.Book
}
