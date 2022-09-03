package userbook

import (
	"github.com/ryota1116/stacked_books/domain/model/book"
)

type UserBookRepository interface {
	CreateOne(userBook UserBook) UserBook
	FindAllByUserId(userId int) ([]book.Book, error)
}
