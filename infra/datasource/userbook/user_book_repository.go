package userbook

import (
	book2 "github.com/ryota1116/stacked_books/domain/model/book"
	"github.com/ryota1116/stacked_books/domain/model/userbook"
	"github.com/ryota1116/stacked_books/infra/datasource"
	"github.com/ryota1116/stacked_books/infra/datasource/book"
	"time"
)

type userBookPersistence struct{}

func NewUserBookPersistence() userbook.UserBookRepository {
	return &userBookPersistence{}
}

type UserBook struct {
	UserId    int `gorm:"primaryKey"`
	BookId    int
	Status    int
	Memo      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Book      book.Book
}

// CreateOne : UserBooksレコードを作成する
func (userBookPersistence) Save(userBook userbook.UserBookInterface) error {
	db := datasource.DbConnect()

	if err := db.Create(&userBook).Error; err != nil {
		return err
	}
	return nil
}

func (ubp userBookPersistence) FindUserBooksByStatus(
	userID int,
	status userbook.StatusInterface,
) ([]userbook.UserBookInterface, error) {
	db := datasource.DbConnect()
	var ubRecords []UserBook

	db.Joins("Book").
		Where("user_books.user_id = ? AND user_books.status = ?", userID, status.Value).
		Find(&ubRecords)

	var bs []userbook.UserBookInterface
	for _, ubRecord := range ubRecords {

		b, err := book2.NewBook(
			&ubRecord.Book.Id,
			ubRecord.Book.GoogleBooksId,
			ubRecord.Book.Title,
			ubRecord.Book.Description,
			nil,
			ubRecord.Book.Isbn10,
			ubRecord.Book.Isbn13,
			ubRecord.Book.PageCount,
			ubRecord.Book.PublishedYear,
			ubRecord.Book.PublishedMonth,
			ubRecord.Book.PublishedDate,
			&ubRecord.Book.CreatedAt,
		)

		ub, err := userbook.NewUserBook(
			ubRecord.BookId,
			ubRecord.UserId,
			ubRecord.Status,
			&ubRecord.Memo,
			b,
		)
		if err != nil {
			return nil, err
		}

		bs = append(bs, ub)
	}
	return bs, nil
}
