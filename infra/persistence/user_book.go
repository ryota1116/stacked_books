package persistence

import (
	"github.com/ryota1116/stacked_books/domain/model"
	"github.com/ryota1116/stacked_books/domain/repository"
	"github.com/ryota1116/stacked_books/handler/http/request/user_book/register_user_books"
)

type userBookPersistence struct {}

func NewUserBookPersistence() repository.UserBookRepository {
	return &userBookPersistence{}
}

// CreateOne : UserBooksレコードを作成する
func (userBookPersistence) CreateOne(userId int, bookId int, requestBody RegisterUserBooks.RequestBody) model.UserBook {
	db := DbConnect()

	userBook := model.UserBook{
		UserId:    userId,
		BookId:    bookId,
		Status:    requestBody.UserBook.Status,
		Memo:      requestBody.UserBook.Memo,
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
