package request

import (
	"github.com/ryota1116/stacked_books/domain/model/googleBooksApi"
	"github.com/ryota1116/stacked_books/handler/http/response"
)

type BookHandlerFormValidator struct {
	GoogleBooksApiRequestBody googleBooksApi.RequestParameter
}

// Validate : BookHandler.SearchBooksのリクエストボディのバリデーション
func (bh BookHandlerFormValidator) Validate() (bool, response.ErrorResponseBody) {
	if bh.GoogleBooksApiRequestBody.Title == "" {
		return false, response.ErrorResponseBody{Message: "本のタイトルを入力してください"}
	}

	return true, response.ErrorResponseBody{Message: ""}
}
