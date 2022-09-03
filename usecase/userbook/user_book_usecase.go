package userbook

import (
	"github.com/ryota1116/stacked_books/domain/model/book"
	"github.com/ryota1116/stacked_books/domain/model/userbook"
)

type UserBookUseCase interface {
	RegisterUserBook(command UserBookCreateCommand) (book.Book, userbook.UserBook)
	FindUserBooksByUserId(userId int) ([]UserBookDto, error)
}

type userBookUseCase struct {
	bookRepository     book.BookRepository
	userBookRepository userbook.UserBookRepository
}

func NewUserBookUseCase(br book.BookRepository, ubr userbook.UserBookRepository) UserBookUseCase {
	return &userBookUseCase{
		bookRepository:     br,
		userBookRepository: ubr,
	}
}

// RegisterUserBook : UserBooksレコードを作成する
func (ubu userBookUseCase) RegisterUserBook(command UserBookCreateCommand) (book.Book, userbook.UserBook) {
	// GoogleBooksIDからBookレコードを検索し、存在しなければ作成する
	b := ubu.bookRepository.FindOrCreateByGoogleBooksId(command.Book.GoogleBooksId)

	userBook := userbook.UserBook{
		UserId: command.UserId,
		BookId: b.Id,
		Status: command.UserBook.Status,
		Memo:   command.UserBook.Memo,
	}

	// UserBooksレコードを作成する
	savedUserBook := ubu.userBookRepository.CreateOne(userBook)

	return b, savedUserBook
}

// FindUserBooksByUserId : ログイン中のユーザーが登録している本の一覧を取得する
func (ubu userBookUseCase) FindUserBooksByUserId(userId int) ([]UserBookDto, error) {
	books, err := ubu.userBookRepository.FindAllByUserId(userId)

	// DTOに変換
	userBooks := DtoGenerator{Books: books}.Execute()

	return userBooks, err
}
