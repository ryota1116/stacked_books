package book

import (
	"github.com/ryota1116/stacked_books/domain/model/book"
	model "github.com/ryota1116/stacked_books/domain/model/searched_books/google_books_api"
)

type BookUseCaseInterface interface {
	SearchBooks(title string) (model.ResponseBodyFromGoogleBooksApi, error)
	GetBookById(bookId int) (BookDto, error)
}

type bookUseCase struct {
	bookRepository       book.BookRepository
	googleBooksAPIClient model.GoogleBooksApiClientInterface
}

func NewBookUseCase(
	br book.BookRepository,
	client model.GoogleBooksApiClientInterface,
) BookUseCaseInterface {
	return &bookUseCase{
		br,
		client,
	}
}

// SearchBooks : 外部APIを用いて書籍検索を行う
func (bu bookUseCase) SearchBooks(title string) (model.ResponseBodyFromGoogleBooksApi, error) {
	// 外部APIで書籍を検索
	// 書籍検索用のレスポンスボディ構造体のスライス型
	responseFromGoogleBooksAPI, err := bu.googleBooksAPIClient.SendRequest(title)
	if err != nil {
		return responseFromGoogleBooksAPI, err
	}

	// GoogleBooksAPIのJSONレスポンスの構造体を返す
	return responseFromGoogleBooksAPI, nil
}

func (bu bookUseCase) GetBookById(bookId int) (BookDto, error) {
	b, err := bu.bookRepository.FindOneById(bookId)
	if err != nil {
		return BookDto{}, err
	}

	bookDtoGenerator := BookDtoGenerator{Book: b}
	bookDto := bookDtoGenerator.Execute()

	return bookDto, nil
}
