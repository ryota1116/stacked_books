package search_books

// RequestBody : 書籍検索用のリクエストボディ構造体
type RequestBody struct {
	Title	string	`json:"title"`
}
