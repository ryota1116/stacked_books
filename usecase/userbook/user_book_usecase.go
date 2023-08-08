package userbook

import (
	model "github.com/ryota1116/stacked_books/domain/model/book"
	"github.com/ryota1116/stacked_books/domain/model/userbook"
	"github.com/ryota1116/stacked_books/usecase/book"
)

type UserBookUseCase interface {
	RegisterUserBook(command UserBookCreateCommand) (book.Dto, UserBookDto, error)
	FindUserBooksByUserId(userId int) ([]book.Dto, error)
	SearchUserBooksByStatus(command SearchUserBooksByStatusCommand) ([]UserBookDto, error)
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
func (ubu userBookUseCase) RegisterUserBook(command UserBookCreateCommand) (book.Dto, UserBookDto, error) {
	// GoogleBooksIDからBookエンティティを取得
	b, err := ubu.bookRepository.FindOneByGoogleBooksId(command.Book.GoogleBooksId)
	if err != nil {
		// 取得できなければ新規作成する
		b, err := model.NewBook(
			nil,
			command.Book.GoogleBooksId,
			command.Book.Title,
			command.Book.Description,
			command.Book.Image,
			command.Book.Isbn10,
			command.Book.Isbn13,
			command.Book.PageCount,
			command.Book.PublishedYear,
			command.Book.PublishedMonth,
			command.Book.PublishedDate,
			nil,
		)

		if err != nil {
			return book.Dto{}, UserBookDto{}, err
		}
		if err := ubu.bookRepository.SaveOne(b); err != nil {
			return book.Dto{}, UserBookDto{}, err
		}
	}

	// UserBookエンティティの生成
	userBook, err := userbook.NewUserBook(
		command.UserId,
		*b.Id().Value(),
		command.UserBook.Status,
		command.UserBook.Memo,
		b,
	)
	if err != nil {
		return book.Dto{}, UserBookDto{}, err
	}

	// UserBookを保存する
	if err := ubu.userBookRepository.SaveOne(userBook); err != nil {
		return book.Dto{}, UserBookDto{}, err
	}

	// DTOを返却
	return book.DtoGenerator{Book: b}.Execute(),
		UserBookDtoGenerator{UserBook: userBook}.Execute(),
		nil
}

// FindUserBooksByUserId : ログイン中のユーザーが登録している本の一覧を取得する
func (ubu userBookUseCase) FindUserBooksByUserId(userId int) ([]book.Dto, error) {
	books, err := ubu.bookRepository.FindListByUserId(userId)

	// DTOに変換
	var booksDto []book.Dto
	for _, b := range books {
		dtog := book.DtoGenerator{Book: b}

		booksDto = append(booksDto, dtog.Execute())
	}

	return booksDto, err
}

func (ubu userBookUseCase) SearchUserBooksByStatus(
	command SearchUserBooksByStatusCommand,
) ([]UserBookDto, error) {
	status, err := userbook.NewStatus(command.Status)
	if err != nil {
		return nil, err
	}

	userBooks, err := ubu.userBookRepository.FindListByStatus(
		command.UserId,
		status,
	)

	// DTOに変換
	var userBooksDto []UserBookDto
	for _, b := range userBooks {
		dtog := UserBookDtoGenerator{UserBook: b}

		userBooksDto = append(userBooksDto, dtog.Execute())
	}
	return userBooksDto, err
}
