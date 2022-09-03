package response

import (
	"encoding/json"
	"net/http"
)

// Response : HTTPレスポンスの構造体
type Response struct {
	StatusCode int
	ResponseBody interface{}
}

// ReturnResponse : クライアントにHTTPレスポンスを返す
func (r Response) ReturnResponse(w http.ResponseWriter) {
	// HTTPレスポンスヘッダーのContent-Typeを設定
	w.Header().Set("Content-Type", "application/json")

	// ステータスコードの設定
	w.WriteHeader(int(r.StatusCode))

	// Structなどの型をJSON文字列に変換して、ResponseWriter(w)に書き込む
	err := json.NewEncoder(w).Encode(r.ResponseBody)

	if err != nil {
		Return500Response(w, err)
	}
}

// Return500Response : 500エラーを返す
func Return500Response(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

	// ステータスコードの設定
	w.WriteHeader(http.StatusInternalServerError)

	// Structなどの型をJSON文字列に変換して、ResponseWriter(w)に書き込む
	errResponse := ErrorResponseBody{Message: err.Error()}
	json.NewEncoder(w).Encode(errResponse)
}
