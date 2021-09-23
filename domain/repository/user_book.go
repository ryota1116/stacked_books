package repository

import (
	"github.com/ryota1116/stacked_books/domain/model"
	"github.com/ryota1116/stacked_books/domain/model/dto"
)

type UserBookRepository interface {
	CreateOne(int, int, dto.RegisterUserBookRequestParameter) model.UserBook
}
