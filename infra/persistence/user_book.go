package persistence

import (
	"fmt"
	"github.com/ryota1116/stacked_books/domain/model"
	"github.com/ryota1116/stacked_books/domain/repository"
)

type userBookPersistence struct {}

func NewUserBookPersistence() repository.UserBookRepository {
	return &userBookPersistence{}
}

func (userBookPersistence) CreateOne(userBookParameter model.UserBookParameter) model.UserBookParameter {
	db := DbConnect()
	db.Model(&model.UserBook{}).Create(map[string]interface{}{
		"UserId": 1,
		"BookId": userBookParameter.BookId,
		"status": userBookParameter.Status,
		"memo": userBookParameter.Memo,
	})

	return userBookParameter
}

// ReadUserBooks : ログイン中のユーザーが登録している本の一覧を取得する
func (userBookPersistence) ReadUserBooks(userId int) model.Book {
	db := DbConnect()
	user := model.User{}
	book := model.Book{}
	//result := db.Debug().First(&user, 1)
	err := db.Model(&user).Association("Books").Find(&book)
	if err != nil {
		fmt.Println("aaa")
	}
	return book
}