package usecase

import (
	"github.com/ryota1116/stacked_books/domain/model"
	"github.com/ryota1116/stacked_books/domain/repository"
)

type UserBookUseCase interface {
	RegisterUserBook(userBookParameter model.UserBookParameter) model.UserBookParameter
	ReadUserBooks(userId int) model.Book
}

type userBookUseCase struct {
	bookRepository repository.BookRepository
	userBookRepository repository.UserBookRepository
}

func NewUserBookUseCase(br repository.BookRepository, ubr repository.UserBookRepository) UserBookUseCase {
	return &userBookUseCase{
		bookRepository:     br,
		userBookRepository: ubr,
	}
}

func (ubu userBookUseCase) RegisterUserBook(userBookParameter model.UserBookParameter) model.UserBookParameter {
	userBookParameter.Book = ubu.bookRepository.FindOrCreateByGoogleBooksId(userBookParameter.GoogleBooksId, userBookParameter)
	userBook := ubu.userBookRepository.CreateOne(userBookParameter)
	return userBook
}

func (ubu userBookUseCase) ReadUserBooks(userId int) model.Book {
	userBooks := ubu.userBookRepository.ReadUserBooks(userId)
	return userBooks
}
