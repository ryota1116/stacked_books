package userbook

import (
	"github.com/ryota1116/stacked_books/domain/model/book"
	"github.com/ryota1116/stacked_books/domain/model/userbook"
	"github.com/ryota1116/stacked_books/handler/http/request/user_book/register_user_books"
	"github.com/ryota1116/stacked_books/infra/persistence"
)

type userBookPersistence struct{}

func NewUserBookPersistence() userbook.UserBookRepository {
	return &userBookPersistence{}
}

// CreateOne : UserBooksレコードを作成する
func (userBookPersistence) CreateOne(userId int, bookId int, requestBody RegisterUserBooks.RequestBody) userbook.UserBook {
	db := persistence.DbConnect()

	userBook := userbook.UserBook{
		UserId: userId,
		BookId: bookId,
		Status: requestBody.UserBook.Status,
		Memo:   requestBody.UserBook.Memo,
	}

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
