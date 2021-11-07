package usecase

import (
	"github.com/ryota1116/stacked_books/domain/model"
	"github.com/ryota1116/stacked_books/domain/model/UserBook"
	"github.com/ryota1116/stacked_books/domain/repository"
)

type UserBookUseCase interface {
	RegisterUserBook(userBookParameter model.UserBookParameter) model.UserBookParameter
	ReadUserBooks(userId int) model.Book
	SearchUserBooksByStatus(userID int, status UserBook.Status) []model.UserBook
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
func (uub userBookUseCase) RegisterUserBook(userId int, registerUserBookRequestParameter dto.RegisterUserBookRequestParameter) dto.RegisterUserBookResponse {
	// GoogleBooksIDからBookレコードを検索し、存在しなければ作成する
	book := uub.bookRepository.FindOrCreateByGoogleBooksId(registerUserBookRequestParameter)
	// UserBooksレコードを作成する
	userBook := uub.userBookRepository.CreateOne(userId, book.Id, registerUserBookRequestParameter)
	// RegisterUserBookResponse構造体を生成する
	userBookResponse := dto.BuildRegisterUserBookResponse(book, userBook)

	return userBookResponse
}

// FindUserBooksByUserId : ログイン中のユーザーが登録している本の一覧を取得する
func (ubu userBookUseCase) FindUserBooksByUserId(userId int) []model.Book {
	userBooks := ubu.userBookRepository.FindAllByUserId(userId)
	return userBooks
}

func (ubu userBookUseCase) SearchUserBooksByStatus(userID int, status UserBook.Status) []model.UserBook {
	return ubu.userBookRepository.FindUserBooksByStatus(userID, status)
}