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
func (userBookPersistence) CreateOne(userBook userbook.UserBook) userbook.UserBook {
	db := datasource.DbConnect()
	db.Create(&userBook)

	return userBook
}
