package userbook

import (
	"github.com/ryota1116/stacked_books/domain/model/userbook"
	"github.com/ryota1116/stacked_books/usecase/book"
	"time"
)

type UserBookDtoGenerator struct {
	UserBook userbook.UserBookInterface
}

type UserBookDto struct {
	UserId    int
	BookId    int
	Status    int
	Memo      *string
	CreatedAt time.Time
	UpdatedAt time.Time
	BookDto   book.BookDto
}

func (dtog UserBookDtoGenerator) Execute() UserBookDto {
	return UserBookDto{
		UserId:    dtog.UserBook.UserId().Value(),
		BookId:    dtog.UserBook.BookId().Value(),
		Status:    dtog.UserBook.Status().Value(),
		Memo:      dtog.UserBook.Memo().Value(),
		CreatedAt: dtog.UserBook.CreatedAt(),
		UpdatedAt: dtog.UserBook.UpdatedAt(),
		BookDto:   book.BookDtoGenerator{Book: dtog.UserBook.Book()}.Execute(),
	}
}
