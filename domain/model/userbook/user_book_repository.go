package userbook

import (
	"github.com/ryota1116/stacked_books/domain/model/book"
	RegisterUserBooks "github.com/ryota1116/stacked_books/handler/http/request/user_book/register_user_books"
)

type UserBookRepository interface {
	CreateOne(int, int, RegisterUserBooks.RequestBody) UserBook
	FindAllByUserId(userId int) ([]book.Book, error)
}
