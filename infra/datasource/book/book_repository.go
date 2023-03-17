package book

import (
	"github.com/ryota1116/stacked_books/domain/model/book"
	"github.com/ryota1116/stacked_books/infra/datasource"
)

type bookPersistence struct{}

func NewBookPersistence() book.BookRepository {
	return &bookPersistence{}
}

func (bookPersistence) FindOneByGoogleBooksId(GoogleBooksId string) (book.Book, error) {
	db := datasource.DbConnect()
	b := book.Book{}

	if err := db.Where("google_books_id = ?", GoogleBooksId).First(&b).Error; err != nil {
		return b, err
	}
	return b, nil
}

func (bookPersistence) Save(book book.Book) error {
	db := datasource.DbConnect()

	if err := db.Create(&book).Error; err != nil {
		return err
	}
	return nil
}

// FindAllByUserId : ログイン中のユーザーが登録している本の一覧を取得する
func (bookPersistence) FindAllByUserId(userId int) ([]book.Book, error) {
	db := datasource.DbConnect()
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
