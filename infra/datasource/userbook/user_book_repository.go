package userbook

import (
	"github.com/ryota1116/stacked_books/domain/model/userbook"
	"github.com/ryota1116/stacked_books/infra/datasource"
)

type userBookPersistence struct{}

func NewUserBookPersistence() userbook.UserBookRepository {
	return &userBookPersistence{}
}

// CreateOne : UserBooksレコードを作成する
func (userBookPersistence) Save(userBook userbook.UserBookInterface) error {
	db := datasource.DbConnect()

	if err := db.Create(&userBook).Error; err != nil {
		return err
	}
	return nil
}
