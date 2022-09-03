package userbook

import (
	"github.com/ryota1116/stacked_books/domain/model/book"
	"github.com/ryota1116/stacked_books/domain/model/userbook"
	"github.com/ryota1116/stacked_books/infra/persistence"
)

type userBookPersistence struct{}

func NewUserBookPersistence() userbook.UserBookRepository {
	return &userBookPersistence{}
}

// CreateOne : UserBooksレコードを作成する
func (userBookPersistence) CreateOne(userBook userbook.UserBook) userbook.UserBook {
	db := persistence.DbConnect()
	db.Create(&userBook)

	return userBook
}

// FindAllByUserId : ログイン中のユーザーが登録している本の一覧を取得する
func (userBookPersistence) FindAllByUserId(userId int) ([]book.Book, error) {
	db := persistence.DbConnect()
	var books []book.Book

	// ユーザーが登録している本一覧を取得
	if err := db.Joins("inner join user_books on books.id = user_books.book_id").
		Joins("inner join users on user_books.user_id = ?", userId).
		Group("books.id").
		Find(&books).Error; err != nil {
		return books, err
	}

	return books, nil
}
