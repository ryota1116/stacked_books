package book

import (
	bookEntity "github.com/ryota1116/stacked_books/domain/model/book"
	"github.com/ryota1116/stacked_books/infra/datasource"
	"time"
)

type bookPersistence struct{}

func NewBookPersistence() bookEntity.BookRepository {
	return &bookPersistence{}
}

type book struct {
	Id             int `gorm:"primaryKey"`
	GoogleBooksId  string
	Title          string
	Description    *string
	Isbn_10        *string
	Isbn_13        *string
	PageCount      int
	PublishedYear  *int
	PublishedMonth *int
	PublishedDate  *int
	CreatedAt      time.Time
}

func (bookPersistence) FindOneByGoogleBooksId(GoogleBooksId string) (bookEntity.BookInterface, error) {
	db := datasource.DbConnect()
	book := book{}

	if err := db.Table("books").
		Where("google_books_id = ?", GoogleBooksId).
		Find(&book).Error; err != nil {
		return nil, err
	}

	u, err := bookEntity.NewBook(
		&book.Id,
		book.GoogleBooksId,
		book.Title,
		book.Description,
		nil,
		book.Isbn_10,
		book.Isbn_13,
		book.PageCount,
		book.PublishedYear,
		book.PublishedMonth,
		book.PublishedDate,
		&book.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (bookPersistence) Save(book bookEntity.BookInterface) error {
	db := datasource.DbConnect()

	if err := db.Create(&book).Error; err != nil {
		return err
	}
	return nil
}

// FindAllByUserId : ログイン中のユーザーが登録している本の一覧を取得する
func (bookPersistence) FindAllByUserId(userId int) ([]bookEntity.BookInterface, error) {
	db := datasource.DbConnect()
	books := []book{}

	// ユーザーが登録している本一覧を取得
	if err := db.
		Joins("inner join user_books on books.id = user_books.book_id").
		Joins("inner join users on user_books.user_id = ?", userId).
		Group("books.id").
		Find(&books).Error; err != nil {
		return nil, err
	}

	bs := []bookEntity.BookInterface{}
	for _, book := range books {
		id := book.Id

		b, err := bookEntity.NewBook(
			&id, // NOTE:　book.Idを直接入れると同じメモリが渡されたので、一時変数idを使っている
			book.GoogleBooksId,
			book.Title,
			book.Description,
			nil,
			book.Isbn_10,
			book.Isbn_13,
			book.PageCount,
			book.PublishedYear,
			book.PublishedMonth,
			book.PublishedDate,
			&book.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		bs = append(bs, b)
	}

	return bs, nil
}
