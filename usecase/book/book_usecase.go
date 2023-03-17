package book

import (
	model "github.com/ryota1116/stacked_books/domain/model/searched_books/google_books_api"
)

type BookUseCaseInterface interface {
	SearchBooks(title string) (model.ResponseBodyFromGoogleBooksApi, error)
}

type bookUseCase struct {
	// BookRepositoryを使う必要が出たときにコメントアウト外す
	// bookRepository repository.BookRepository
	googleBooksAPIClient model.GoogleBooksApiClientInterface
}

func NewBookUseCase(client model.GoogleBooksApiClientInterface) BookUseCaseInterface {
	return &bookUseCase{
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
