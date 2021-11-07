package persistence

import (
	"fmt"
	"github.com/ryota1116/stacked_books/domain/model"
	"github.com/ryota1116/stacked_books/domain/model/UserBook"
	"github.com/ryota1116/stacked_books/domain/repository"
)

type userBookPersistence struct {}

func NewUserBookPersistence() repository.UserBookRepository {
	return &userBookPersistence{}
}

func (userBookPersistence) CreateOne(userBookParameter model.UserBookParameter) model.UserBookParameter {
	db := DbConnect()
	db.Model(&model.UserBook{}).Create(map[string]interface{}{
		"UserId": userBookParameter.UserId,
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

	// ユーザーを取得する
	db.Where("id = ?", userId).First(&user)
	// ユーザーが登録している本一覧を取得
	err := db.Model(&user).Association("Books").Find(&book)

	fmt.Println()
	if err != nil {
		fmt.Println("aaa")
	}

	return book
}

// FindUserBooksByStatus : 読書ステータスでユーザーが登録している本一覧を取得する
func (userBookPersistence) FindUserBooksByStatus(userID int, status UserBook.Status) []model.UserBook {
	db := DbConnect()
	var userBooks []model.UserBook

	db.Joins("Book").
		Where("user_books.user_id = ? AND user_books.status = ?", userID, status.Value).
		Find(&userBooks)

	return userBooks
}
