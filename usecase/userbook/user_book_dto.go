package userbook

import (
	"github.com/ryota1116/stacked_books/domain/model/userbook"
	"time"
)

type UserBookDtoGenerator struct {
	UserBook userbook.UserBook
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
		Id:        dtog.UserBook.Id,
		UserId:    dtog.UserBook.UserId,
		BookId:    dtog.UserBook.BookId,
		Status:    dtog.UserBook.Status,
		Memo:      dtog.UserBook.Memo,
		CreatedAt: dtog.UserBook.CreatedAt,
		UpdatedAt: dtog.UserBook.UpdatedAt,
	}
}
