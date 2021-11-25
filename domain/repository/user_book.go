package repository

import (
	"github.com/ryota1116/stacked_books/domain/model"
	"github.com/ryota1116/stacked_books/domain/model/dto"
)

type UserBookRepository interface {
	CreateOne(userId int, bookId int, registerUserBookRequestParameter dto.RegisterUserBookRequestParameter) (model.UserBook, error)
}
