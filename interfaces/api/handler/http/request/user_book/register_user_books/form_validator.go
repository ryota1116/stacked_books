package RegisterUserBooks

import (
	"fmt"
	"github.com/ryota1116/stacked_books/domain/model/userbook"
	"github.com/ryota1116/stacked_books/interfaces/api/handler/http/response"
	"unicode/utf8"
)

const maxMemoWordCount = 255

type FormValidator struct {
	RequestBody
}

// Validate : UserBookHandler.RegisterUserBooksのリクエストボディのバリデーション
func (rbh FormValidator) Validate() (bool, response.ErrorResponseBody) {
	if rbh.RequestBody.Book.GoogleBooksId == "" {
		return false, response.ErrorResponseBody{Message: "GoogleBooksIdが空になっています。"}
	}

	if rbh.RequestBody.Book.Title == "" {
		return false, response.ErrorResponseBody{Message: "本のタイトルを入力してください。"}
	}

	if rbh.RequestBody.Book.Description == "" {
		return false, response.ErrorResponseBody{Message: "本の説明文を入力してください。"}
	}

	if rbh.RequestBody.Book.PageCount <= 0 {
		return false, response.ErrorResponseBody{Message: "本のページ数は1ページ以上で入力してください。"}
	}

	// Contain関数を作成する https://zenn.dev/glassonion1/articles/7c7830a269909c
	isValidStatus := func() bool {
		for _, bookStatus := range userbook.GetBookStatuses() {
			if rbh.RequestBody.UserBook.Status == bookStatus {
				return true
			}
		}
		return false
	}

	if !isValidStatus() {
		return isValidStatus(), response.ErrorResponseBody{Message: fmt.Sprintf(
			"読書ステータスの値が不正です。 status: %d",
			rbh.RequestBody.UserBook.Status),
		}
	}

	memoCount := utf8.RuneCountInString(rbh.RequestBody.UserBook.Memo)
	if memoCount > maxMemoWordCount {
		return false, response.ErrorResponseBody{Message: "メモは255文字以下で入力ください。"}
	}

	return true, response.ErrorResponseBody{Message: ""}
}
