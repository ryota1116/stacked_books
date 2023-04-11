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
	Book() book.BookInterface

	ChangeMemo(value *string) error
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
	book book.BookInterface,
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
		book:   book,
	}, nil
}

func (ub *userBook) UserId() UserIdInterface {
	return ub.userId
}

func (ub *userBook) BookId() BookIdInterface {
	return ub.bookId
}

func (ub *userBook) Status() StatusInterface {
	return ub.status
}

func (ub *userBook) Memo() MemoInterface {
	return ub.memo
}

func (ub *userBook) CreatedAt() time.Time {
	return ub.createdAt
}

func (ub *userBook) UpdatedAt() time.Time {
	return ub.updatedAt
}

func (ub *userBook) Book() book.BookInterface {
	return ub.book
}

func (ub *userBook) ChangeMemo(value *string) error {
	return ub.memo.changeMemo(value)
}
