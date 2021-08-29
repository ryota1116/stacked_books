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
		"BookId": userBookParameter.Book.Id,
		"status": userBookParameter.Status,
		"memo": userBookParameter.Memo,
	})

	return userBookParameter
}

// ReadUserBooks : ログイン中のユーザーが登録している本の一覧を取得する
func (userBookPersistence) ReadUserBooks(userId int) []model.Book {
	db := DbConnect()
	user := model.User{}
	var books []model.Book

	// ユーザーを取得する
	db.Where("id = ?", userId).First(&user)
	// ユーザーが登録している本一覧を取得
	err := db.Model(&user).Association("Books").Find(&books)
	fmt.Println("---------")
	fmt.Println(books)

	fmt.Println()
	if err != nil {
		fmt.Println("aaa")
	}

	return books
}

func (userBookPersistence) FindUserBooksWithReadingStatus(userId int, readingStatus int) []model.Book {
	db := DbConnect()
	var books []model.Book

	// 特定の読書ステータスに紐付くBooks一覧を取得
	err := db.Find(&books).
		Joins(
			"JOIN user_books ON user_books.user_id = users.id AND user_books.status = ?",
			readingStatus).
		Where("users.id = ?", userId)

	if err != nil {
		fmt.Println("aaa")
	}

	return books
}
