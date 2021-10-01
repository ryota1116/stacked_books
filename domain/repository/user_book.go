package repository

import (
	"github.com/ryota1116/stacked_books/domain/model"
	"google.golang.org/grpc/status"
)

type UserBookRepository interface {
	CreateOne(userBookParameter model.UserBookParameter) model.UserBookParameter
	ReadUserBooks(userId int) model.Book
	SearchUserBooksByStatus(status.Status)
}
