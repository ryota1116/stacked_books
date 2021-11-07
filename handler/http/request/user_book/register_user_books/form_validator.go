package RegisterUserBooks

import (
	"github.com/ryota1116/stacked_books/domain/model/UserBook"
	"github.com/ryota1116/stacked_books/domain/model/dto"
	"github.com/ryota1116/stacked_books/handler/http/response"
	"unicode/utf8"
)

const maxMemoWordCount = 255

type FormValidator struct {
	RegisterUserBookRequestParameter dto.RegisterUserBookRequestParameter
}

// Validate : UserBookHandler.RegisterUserBooksのリクエストボディのバリデーション
func (rbh FormValidator) Validate() (bool, response.ErrorResponseBody) {
	if rbh.RegisterUserBookRequestParameter.Book.GoogleBooksId == "" {
		return false, response.ErrorResponseBody{Message: "GoogleBooksIdが空になっています。"}
	}

	if rbh.RegisterUserBookRequestParameter.Book.Title == "" {
		return false, response.ErrorResponseBody{Message: "本のタイトルを入力してください。"}
	}

	if rbh.RegisterUserBookRequestParameter.Book.Description == "" {
		return false, response.ErrorResponseBody{Message: "本の説明文を入力してください。"}
	}

	if rbh.RegisterUserBookRequestParameter.Book.PageCount <= 0 {
		return false, response.ErrorResponseBody{Message: "本のページ数は1ページ以上で入力してください。"}
	}

	for _, bookStatus := range UserBook.GetBookStatuses() {
		if rbh.RegisterUserBookRequestParameter.UserBook.Status == bookStatus {
			continue
		}
		return false, response.ErrorResponseBody{Message: "読書ステータスの値が不正です。"}
	}

	memoCount := utf8.RuneCountInString(rbh.RegisterUserBookRequestParameter.UserBook.Memo)
	if memoCount > maxMemoWordCount {
		return false, response.ErrorResponseBody{Message: "メモは255文字以下で入力ください。"}
	}

	return true, response.ErrorResponseBody{Message: ""}
}