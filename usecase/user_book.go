package usecase

import (
	"github.com/ryota1116/stacked_books/domain/model"
	"github.com/ryota1116/stacked_books/domain/repository"
)

type UserBookUseCase interface {
	RegisterUserBook(userId int, userBookParameter model.UserBookParameter) model.UserBookParameter
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

func (uub userBookUseCase) RegisterUserBook(userId int, userBookParameter model.UserBookParameter) model.UserBookParameter {
	userBookParameter.Book = uub.bookRepository.FindOrCreateByGoogleBooksId(userBookParameter.GoogleBooksId, userBookParameter)
	userBook := uub.userBookRepository.CreateOne(userId, userBookParameter)
	return userBook
}
