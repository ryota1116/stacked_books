package usecase

import (
	"github.com/ryota1116/stacked_books/domain/model"
	"github.com/ryota1116/stacked_books/domain/repository"
	"github.com/ryota1116/stacked_books/handler/http/request/user_book/register_user_books"
)

type UserBookUseCase interface {
	RegisterUserBook(int, RegisterUserBooks.RequestBody) (model.Book, model.UserBook)
	FindUserBooksByUserId(userId int) ([]model.Book, error)
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
func (uub userBookUseCase) RegisterUserBook(userId int, requestBody RegisterUserBooks.RequestBody) (model.Book, model.UserBook) {
	// GoogleBooksIDからBookレコードを検索し、存在しなければ作成する
	book := uub.bookRepository.FindOrCreateByGoogleBooksId(requestBody)
	// UserBooksレコードを作成する
	userBook := uub.userBookRepository.CreateOne(userId, book.Id, requestBody)

	return book, userBook
}

// FindUserBooksByUserId : ログイン中のユーザーが登録している本の一覧を取得する
func (ubu userBookUseCase) FindUserBooksByUserId(userId int) ([]model.Book, error) {
	userBooks, err := ubu.userBookRepository.FindAllByUserId(userId)
	return userBooks, err
}
