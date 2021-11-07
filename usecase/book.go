package usecase

import (
	"github.com/ryota1116/stacked_books/domain/model/googleBooksApi"
)

type BookUseCaseInterface interface {
	SearchBooks(requestParameter googleBooksApi.RequestParameter) (googleBooksApi.SearchBooksResponses, error)
}

type bookUseCase struct {
	// BookRepositoryを使う必要が出たときにコメントアウト外す
	// bookRepository repository.BookRepository
}

func NewBookUseCase() BookUseCaseInterface {
	return &bookUseCase{}
}

// SearchBooks : 外部APIを用いて書籍検索を行う
func (bu bookUseCase) SearchBooks(requestParameter googleBooksApi.RequestParameter) (googleBooksApi.SearchBooksResponses, error) {
	// 外部APIで書籍を検索
	// 書籍検索用のレスポンスボディ構造体のスライス型
	return googleBooksApi.Client{}.SendRequest(requestParameter.Title)
}
