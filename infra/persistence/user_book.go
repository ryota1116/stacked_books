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
func (userBookPersistence) FindAllByUserId(userId int) ([]model.Book, error){
	db := DbConnect()
	var books []model.Book

	// ユーザーが登録している本一覧を取得
	if err := db.Joins("inner join user_books on books.id = user_books.book_id").
		Joins("inner join users on user_books.user_id = ?", userId).
		Group("books.id").
		Find(&books).Error; err != nil {
		return books, err
	}

	return books, nil
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
