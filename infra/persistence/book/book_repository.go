package book

import (
	"github.com/ryota1116/stacked_books/domain/model/book"
	"github.com/ryota1116/stacked_books/infra/persistence"
)

type bookPersistence struct{}

func NewBookPersistence() book.BookRepository {
	return &bookPersistence{}
}

// FindOrCreateByGoogleBooksId : GoogleBooksIDからBookレコードを検索し、存在しなければ作成する
func (bookPersistence) FindOrCreateByGoogleBooksId(GoogleBooksId string) book.Book {
	db := persistence.DbConnect()
	b := book.Book{}
	db.Where("google_books_id = ?", GoogleBooksId).FirstOrCreate(&b)

	return b
}

// FindAllByUserId : ログイン中のユーザーが登録している本の一覧を取得する
func (bookPersistence) FindAllByUserId(userId int) ([]book.Book, error) {
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
