package repository

import (
	"github.com/ryota1116/stacked_books/domain/model"
	RegisterUserBooks "github.com/ryota1116/stacked_books/handler/http/request/user_book/register_user_books"
)

type UserBookRepository interface {
	CreateOne(int, int, RegisterUserBooks.RequestBody) model.UserBook
	FindAllByUserId(userId int) ([]model.Book, error)
	FindUserBooksWithReadingStatus(userId int, readingStatus int) ([]model.Book, error)
}
