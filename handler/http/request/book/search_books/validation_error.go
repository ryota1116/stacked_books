package search_books

// ValidationError : リクエストのバリデーションエラー
type ValidationError struct {
	Message string `json:"message"`
}
