package repository

import (
	"github.com/ryota1116/stacked_books/domain/model"
	"github.com/ryota1116/stacked_books/domain/model/dto"
)

type BookRepository interface {
	FindOrCreateByGoogleBooksId(parameter dto.RegisterUserBookRequestParameter) model.Book
}
