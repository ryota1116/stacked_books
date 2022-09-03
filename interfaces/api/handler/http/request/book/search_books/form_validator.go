package search_books

type FormValidator struct {
	GoogleBooksApiRequestParameter RequestParameter
}

// Validate : BookHandler.SearchBooksのリクエストパラメーターのバリデーション
func (bh FormValidator) Validate() (bool, ValidationError) {
	if bh.GoogleBooksApiRequestParameter.Title == "" {
		return false, ValidationError{Message: "本のタイトルを入力してください"}
	}

	return true, ValidationError{Message: ""}
}
