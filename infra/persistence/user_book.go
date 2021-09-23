package persistence

import (
	"fmt"
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

// FindAllByUserId : ログイン中のユーザーが登録している本の一覧を取得する
func (userBookPersistence) FindAllByUserId(userId int) model.Book {
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
