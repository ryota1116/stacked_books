package persistence

import (
	"github.com/ryota1116/stacked_books/domain/model"
	"github.com/ryota1116/stacked_books/domain/repository"
)

type userBookPersistence struct {}

func NewUserBookPersistence() repository.UserBookRepository {
	return &userBookPersistence{}
}

func (userBookPersistence) CreateOne(userId int, userBookParameter model.UserBookParameter) model.UserBookParameter {
	db := DbConnect()

	userBook := model.UserBook{
		UserId:    userId,
		BookId:    userBookParameter.Book.Id,
		Status:    userBookParameter.Status,
		Memo:      userBookParameter.Memo,
	}

	db.Create(&userBook)

	return userBookParameter
}
