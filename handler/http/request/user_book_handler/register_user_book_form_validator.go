package user_book_handler

import (
	"github.com/ryota1116/stacked_books/domain/model/dto"
	"github.com/ryota1116/stacked_books/handler/http/response"
	"unicode/utf8"
)

const maxMemoWordCount = 255

type RegisterBookFormValidator struct {
	RegisterUserBookRequestParameter dto.RegisterUserBookRequestParameter
}

// Validate : UserBookHandler.RegisterUserBooksのリクエストボディのバリデーション
func (rbh RegisterBookFormValidator) Validate() (bool, response.ErrorResponseBody) {
	if rbh.RegisterUserBookRequestParameter.Book.GoogleBooksId == "" {
		return false, response.ErrorResponseBody{Message: "本のタイトルを入力してください"}
	}

	if rbh.RegisterUserBookRequestParameter.Book.Title == "" {
		return false, response.ErrorResponseBody{Message: "本のタイトルを入力してください"}
	}

	if rbh.RegisterUserBookRequestParameter.Book.Description == "" {
		return false, response.ErrorResponseBody{Message: "本のタイトルを入力してください"}
	}

	if rbh.RegisterUserBookRequestParameter.Book.PageCount == "" {
		return false, response.ErrorResponseBody{Message: "本のタイトルを入力してください"}
	}

	if rbh.RegisterUserBookRequestParameter.UserBook.Status == "" {
		return false, response.ErrorResponseBody{Message: "本のタイトルを入力してください"}
	}

	memoCount := utf8.RuneCountInString(rbh.RegisterUserBookRequestParameter.UserBook.Memo)
	if memoCount > maxMemoWordCount {
		return false, response.ErrorResponseBody{Message: "メモは255文字以下で入力ください。"}
	}

	return true, response.ErrorResponseBody{Message: ""}
}