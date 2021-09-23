package persistence

import (
	"github.com/ryota1116/stacked_books/domain/model"
	"github.com/ryota1116/stacked_books/domain/model/dto"
	"github.com/ryota1116/stacked_books/domain/repository"
)

type userBookPersistence struct {}

func NewUserBookPersistence() repository.UserBookRepository {
	return &userBookPersistence{}
}

// CreateOne : UserBooksレコードを作成する
func (userBookPersistence) CreateOne(userId int, bookId int, registerUserBookRequestParameter dto.RegisterUserBookRequestParameter) model.UserBook {
	db := DbConnect()

	userBook := model.UserBook{
		UserId:    userId,
		BookId:    bookId,
		Status:    registerUserBookRequestParameter.UserBook.Status,
		Memo:      registerUserBookRequestParameter.UserBook.Memo,
	}

	db.Create(&userBook)

	return userBook
}
