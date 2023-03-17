package userbook

import (
	"github.com/ryota1116/stacked_books/domain/model/book"
	userBookUseCase "github.com/ryota1116/stacked_books/usecase/userbook"
	"time"
)

type UserBook struct {
	Id        int
	UserId    int
	BookId    int
	Status    Status
	Memo      Memo
	CreatedAt time.Time
	UpdatedAt time.Time
	Book      book.Book
}

// NewUserBook : コンストラクター
func NewUserBook(command userBookUseCase.UserBookCreateCommand, book book.Book) (UserBook, error) {
	status, err := NewStatus(command.UserBook.Status)
	if err != nil {
		return UserBook{}, err
	}

	memo, err := NewMemo(command.UserBook.Memo)
	if err != nil {
		return UserBook{}, err
	}

	return UserBook{
		UserId: command.UserId,
		BookId: book.Id,
		Status: status,
		Memo:   memo,
	}, nil
}