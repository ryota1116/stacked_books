package SearchUserBooksByStatus

import (
	"github.com/ryota1116/stacked_books/domain/model/UserBook"
	"github.com/ryota1116/stacked_books/handler/http/response"
)

type FormValidator struct {
	SearchUserBooksByStatusRequestBody RequestBody
}

// Validate : userBookHandler.SearchUserBooksByStatusのリクエストボディのバリデーション
func (req FormValidator) Validate() (bool, response.ErrorResponseBody) {
	for _, bookStatus := range UserBook.GetBookStatuses() {
		if req.SearchUserBooksByStatusRequestBody.Status == bookStatus {
			return true, response.ErrorResponseBody{Message: ""}
		}
	}

	return false, response.ErrorResponseBody{Message: "読書ステータスの値が不正です。"}
}
