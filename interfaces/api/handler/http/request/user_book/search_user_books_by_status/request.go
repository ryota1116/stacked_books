package SearchUserBooksByStatus

// RequestBody 読書ステータスでユーザーが登録している本一覧を取得するAPIのリクエストボディ構造体
type RequestBody struct {
	Status int `json:"status"`
}
