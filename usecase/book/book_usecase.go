package book

import (
	"github.com/ryota1116/stacked_books/domain/model/google-books-api"
)

type BookUseCaseInterface interface {
	SearchBooks(title string) (google_books_api.ResponseBodyFromGoogleBooksAPI, error)
}

type bookUseCase struct {
	// BookRepositoryを使う必要が出たときにコメントアウト外す
	// bookRepository repository.BookRepository
	googleBooksAPIClient google_books_api.IGoogleBooksAPIClient
}

func NewBookUseCase(client google_books_api.IGoogleBooksAPIClient) BookUseCaseInterface {
	return &bookUseCase{
		client,
	}
}

// SearchBooks : 外部APIを用いて書籍検索を行う
func (bu bookUseCase) SearchBooks(title string) (google_books_api.ResponseBodyFromGoogleBooksAPI, error) {
	// 外部APIで書籍を検索
	// 書籍検索用のレスポンスボディ構造体のスライス型
	responseFromGoogleBooksAPI, err := bu.googleBooksAPIClient.SendRequest(title)
	if err != nil {
		return responseFromGoogleBooksAPI, err
	}

	// GoogleBooksAPIのJSONレスポンスの構造体を返す
	return responseFromGoogleBooksAPI, nil
}
