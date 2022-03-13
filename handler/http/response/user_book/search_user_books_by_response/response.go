package searchUserBooksByResponse

import "time"

type ResponseGenerator struct {

}

// ResponseBody : 読書ステータスでユーザーが登録している本一覧を取得するAPIのレスポンスボディ構造体
type ResponseBody []searchUserBooksByResponse

type searchUserBooksByResponse struct {
	Book struct {
		GoogleBooksId string    `json:"google_books_id"`
		Title         string    `json:"title"`
		Description   string    `json:"description"`
		Image         string    `json:"image"`
	} `json:"book"`
	UserBook struct {
		Status        int       `json:"status"`
		Memo          string    `json:"memo"`
		CreatedAt     time.Time `json:"created_at"`
	} `json:"user_book"`
}

