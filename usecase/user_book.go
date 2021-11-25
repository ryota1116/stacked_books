package usecase

import (
	"github.com/ryota1116/stacked_books/domain/model"
	"github.com/ryota1116/stacked_books/domain/model/dto"
	"github.com/ryota1116/stacked_books/domain/repository"
)

type UserBookUseCase interface {
	RegisterUserBook(int, dto.RegisterUserBookRequestParameter) (model.Book, model.UserBook, error)
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

// RegisterUserBook : UserBooksレコードを作成する
func (uub userBookUseCase) RegisterUserBook(userId int, registerUserBookRequestParameter dto.RegisterUserBookRequestParameter) (model.Book, model.UserBook, error) {
	// GoogleBooksIDからBookレコードを検索し、存在しなければ作成する
	book := uub.bookRepository.FindOrCreateByGoogleBooksId(registerUserBookRequestParameter)
	// UserBooksレコードを作成する
	userBook, err := uub.userBookRepository.CreateOne(userId, book.Id, registerUserBookRequestParameter)

	return book, userBook, err
}
