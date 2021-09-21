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
	db.Model(&model.UserBook{}).Create(map[string]interface{}{
		"UserId": userId,
		"BookId": userBookParameter.Book.Id,
		"status": userBookParameter.Status,
		"memo": userBookParameter.Memo,
	})

	return userBookParameter
}