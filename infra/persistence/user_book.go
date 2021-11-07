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

// CreateOne : UserBooksレコードを作成する
func (userBookPersistence) CreateOne(userId int, bookId int, registerUserBookRequestParameter dto.RegisterUserBookRequestParameter) model.UserBook {
	db := DbConnect()
	db.Model(&model.UserBook{}).Create(map[string]interface{}{
		"UserId": userBookParameter.UserId,
		"BookId": userBookParameter.BookId,
		"status": userBookParameter.Status,
		"memo": userBookParameter.Memo,
	})

	return userBookParameter
}

// FindAllByUserId : ログイン中のユーザーが登録している本の一覧を取得する
func (userBookPersistence) FindAllByUserId(userId int) []model.Book {
	db := DbConnect()
	var books []model.Book

	// ユーザーが登録している本一覧を取得
	err := db.Joins("inner join user_books on books.id = user_books.book_id").
		Joins("inner join users on user_books.user_id = ?", userId).
		Group("books.id").
		Find(&books)

	if err != nil {
		fmt.Println("aaa")
	}

	return books
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
