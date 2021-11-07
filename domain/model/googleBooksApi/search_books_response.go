package googleBooksApi

type SearchBooksResponses []SearchBooksResponse

// SearchBooksResponse : 書籍検索用のレスポンスボディ構造体
//GoogleBooksAPIを叩いた時に取得したJSONレスポンスのうち、
// 必要なフィールドだけをセットしたレスポンスボディの構造体
type SearchBooksResponse struct {
	GoogleBooksId string    `json:"google_books_id"`
	Title	string			`json:"title"`
	Authors	[]string		`json:"authors"`
	Description	string		`json:"description"`
	Isbn10 string			`json:"isbn_10"`
	Isbn13 string			`json:"isbn_13"`
	PageCount int 			`json:"page_count"`
	RegisteredAt string	    `json:"created_at"`
}
