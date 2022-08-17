package book

import (
	"github.com/ryota1116/stacked_books/domain/model/googleBooksApi"
	"github.com/ryota1116/stacked_books/handler/http/request/book/search_books"
)

type BookUseCaseInterface interface {
	SearchBooks(requestParameter search_books.RequestParameter) (googleBooksApi.ResponseBodyFromGoogleBooksAPI, error)
}

type bookUseCase struct {
	// BookRepositoryを使う必要が出たときにコメントアウト外す
	// bookRepository repository.BookRepository
	googleBooksAPIClient googleBooksApi.IGoogleBooksAPIClient
}

func NewBookUseCase(client googleBooksApi.IGoogleBooksAPIClient) BookUseCaseInterface {
	return &bookUseCase{
		client,
	}
}

// SearchBooks : 外部APIを用いて書籍検索を行う
func (bu bookUseCase) SearchBooks(requestParameter search_books.RequestParameter) (googleBooksApi.ResponseBodyFromGoogleBooksAPI, error) {
	// 外部APIで書籍を検索
	// 書籍検索用のレスポンスボディ構造体のスライス型
	responseFromGoogleBooksAPI, err := bu.googleBooksAPIClient.SendRequest(requestParameter.Title)
	if err != nil {
		return responseFromGoogleBooksAPI, err
	}

	// GoogleBooksAPIのJSONレスポンスの構造体を返す
	return responseFromGoogleBooksAPI, nil
}
