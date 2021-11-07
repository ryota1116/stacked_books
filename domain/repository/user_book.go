package repository

import (
	"github.com/ryota1116/stacked_books/domain/model"
	"github.com/ryota1116/stacked_books/domain/model/UserBook"
)

type UserBookRepository interface {
	CreateOne(userBookParameter model.UserBookParameter) model.UserBookParameter
	ReadUserBooks(userId int) model.Book
	FindUserBooksByStatus(userID int, status UserBook.Status) []model.UserBook
}
