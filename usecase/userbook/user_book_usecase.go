package userbook

import (
	model "github.com/ryota1116/stacked_books/domain/model/book"
	"github.com/ryota1116/stacked_books/domain/model/userbook"
	"github.com/ryota1116/stacked_books/usecase/book"
)

type UserBookUseCase interface {
	RegisterUserBook(command UserBookCreateCommand) (book.BookDto, UserBookDto, error)
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
func (ubu userBookUseCase) RegisterUserBook(command UserBookCreateCommand) (book.BookDto, UserBookDto, error) {
	// GoogleBooksIDからBookレコードを検索し、存在しなければ作成する
	var b model.Book
	b, err := ubu.bookRepository.FindOneByGoogleBooksId(command.Book.GoogleBooksId)
	if err != nil {
		b = model.Book{
			GoogleBooksId:  command.Book.GoogleBooksId,
			Title:          command.Book.Title,
			Description:    command.Book.Description,
			Isbn_10:        command.Book.Isbn10,
			Isbn_13:        command.Book.Isbn13,
			PageCount:      command.Book.PageCount,
			PublishedYear:  command.Book.PublishedYear,
			PublishedMonth: command.Book.PublishedMonth,
			PublishedDate:  command.Book.PublishedDate,
		}
		if err := ubu.bookRepository.Save(b); err != nil {
			return book.BookDto{}, UserBookDto{}, err
		}
	}

	userBook, err := userbook.NewUserBook(command, b)
	if err != nil {
		return book.BookDto{}, UserBookDto{}, err
	}

	// UserBooksレコードを作成する
	if err := ubu.userBookRepository.Save(userBook); err != nil {
		return book.BookDto{}, UserBookDto{}, err
	}

	return book.BookDtoGenerator{Book: b}.Execute(),
		UserBookDtoGenerator{UserBook: userBook}.Execute(),
		nil
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
