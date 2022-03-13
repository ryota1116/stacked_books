package search_books

type FormValidator struct {
	GoogleBooksApiRequestBody RequestBody
}

// Validate : BookHandler.SearchBooksのリクエストボディのバリデーション
func (bh FormValidator) Validate() (bool, ValidationError) {
	if bh.GoogleBooksApiRequestBody.Title == "" {
		return false, ValidationError{Message: "本のタイトルを入力してください"}
	}

	return true, ValidationError{Message: ""}
}
