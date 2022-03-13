package repository

import (
	"github.com/ryota1116/stacked_books/domain/model"
	"github.com/ryota1116/stacked_books/domain/model/UserBook"
	RegisterUserBooks "github.com/ryota1116/stacked_books/handler/http/request/user_book/register_user_books"
)

type UserBookRepository interface {
	CreateOne(int, int, RegisterUserBooks.RequestBody) model.UserBook
	FindAllByUserId(userId int) ([]model.Book, error)
	FindUserBooksByStatus(userID int, status UserBook.Status) []model.UserBook
}
