package userbook

import (
	"github.com/ryota1116/stacked_books/domain/model/userbook"
	"time"
)

type UserBookDtoGenerator struct {
	UserBook userbook.UserBookInterface
}

type UserBookDto struct {
	Id        int
	UserId    int
	BookId    int
	Status    int
	Memo      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (dtog UserBookDtoGenerator) Execute() UserBookDto {
	return UserBookDto{
		Id:        dtog.UserBook.Id().Value(),
		UserId:    dtog.UserBook.UserId().Value(),
		BookId:    dtog.UserBook.BookId().Value(),
		Status:    dtog.UserBook.Status().Value(),
		Memo:      dtog.UserBook.Memo().Value(),
		CreatedAt: dtog.UserBook.CreatedAt(),
		UpdatedAt: dtog.UserBook.UpdatedAt(),
	}
}
