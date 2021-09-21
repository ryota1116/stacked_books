package repository

import "github.com/ryota1116/stacked_books/domain/model"

type UserBookRepository interface {
	CreateOne(userId int, userBookParameter model.UserBookParameter) model.UserBookParameter
}
