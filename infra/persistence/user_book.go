package persistence

import (
	"github.com/ryota1116/stacked_books/domain/model"
	"github.com/ryota1116/stacked_books/domain/model/dto"
	"github.com/ryota1116/stacked_books/domain/repository"
	"github.com/ryota1116/stacked_books/infra"
)

type userBookPersistence struct {}

func NewUserBookPersistence() repository.UserBookRepository {
	return &userBookPersistence{}
}

// CreateOne : UserBooksレコードを作成する
func (userBookPersistence) CreateOne(userId int, bookId int, registerUserBookRequestParameter dto.RegisterUserBookRequestParameter) (model.UserBook, error) {
	userBook := model.UserBook{
		UserId:    userId,
		BookId:    bookId,
		Status:    registerUserBookRequestParameter.UserBook.Status,
		Memo:      registerUserBookRequestParameter.UserBook.Memo,
	}

	// Link : https://gorm.io/ja_JP/docs/error_handling.html
	if err := infra.Db.Create(&userBook).Error; err != nil {
		return userBook, err
	}
	return userBook, nil
}
