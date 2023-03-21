package userbook

import (
	"github.com/ryota1116/stacked_books/domain/model/book"
	"time"
)

type UserBookInterface interface {
	UserId() UserIdInterface
	BookId() BookIdInterface
	Status() StatusInterface
	Memo() MemoInterface
	CreatedAt() time.Time
	UpdatedAt() time.Time
}

type userBook struct {
	userId    UserIdInterface
	bookId    BookIdInterface
	status    StatusInterface
	memo      MemoInterface
	createdAt time.Time
	updatedAt time.Time
	book      book.BookInterface
}

func NewUserBook(
	userId int,
	bookId int,
	status int,
	memo *string,
) (UserBookInterface, error) {
	s, err := NewStatus(status)
	if err != nil {
		return &userBook{}, err
	}

	m, err := NewMemo(memo)
	if err != nil {
		return &userBook{}, err
	}

	return &userBook{
		userId: NewUserId(userId),
		bookId: NewBookId(bookId),
		status: s,
		memo:   m,
	}, nil
}

func (u *userBook) UserId() UserIdInterface {
	return u.userId
}

func (u *userBook) BookId() BookIdInterface {
	return u.bookId
}

func (u *userBook) Status() StatusInterface {
	return u.status
}

func (u *userBook) Memo() MemoInterface {
	return u.memo
}

func (u *userBook) CreatedAt() time.Time {
	return u.createdAt
}

func (u *userBook) UpdatedAt() time.Time {
	return u.updatedAt
}
