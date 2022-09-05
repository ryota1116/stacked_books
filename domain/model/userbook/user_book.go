package userbook

import (
	"github.com/ryota1116/stacked_books/domain/model/book"
	"time"
)

type UserBook struct {
	Id        int
	UserId    int
	BookId    int
	Status    int
	Memo      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Book      book.Book
}
