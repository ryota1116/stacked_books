package userbook

import (
	model "github.com/ryota1116/stacked_books/domain/model/book"
	"github.com/ryota1116/stacked_books/domain/model/userbook"
	"github.com/ryota1116/stacked_books/usecase/book"
)

type UserBookUseCase interface {
	RegisterUserBook(command UserBookCreateCommand) (book.BookDto, UserBookDto)
	FindUserBooksByUserId(userId int) ([]book.BookDto, error)
}

type userBookUseCase struct {
	bookRepository     model.BookRepository
	userBookRepository userbook.UserBookRepository
}

func NewUserBookUseCase(br model.BookRepository, ubr userbook.UserBookRepository) UserBookUseCase {
	return &userBookUseCase{
		bookRepository:     br,
		userBookRepository: ubr,
	}
}

// RegisterUserBook : UserBooksレコードを作成する
func (ubu userBookUseCase) RegisterUserBook(command UserBookCreateCommand) (book.BookDto, UserBookDto) {
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

	return book.BookDtoGenerator{Book: b}.Execute(),
		UserBookDtoGenerator{UserBook: savedUserBook}.Execute()
}

// FindUserBooksByUserId : ログイン中のユーザーが登録している本の一覧を取得する
func (ubu userBookUseCase) FindUserBooksByUserId(userId int) ([]book.BookDto, error) {
	books, err := ubu.bookRepository.FindAllByUserId(userId)

	// DTOに変換
	var booksDto []book.BookDto
	for _, b := range books {
		dtog := book.BookDtoGenerator{Book: b}

		booksDto = append(booksDto, dtog.Execute())
	}

	return booksDto, err
}
