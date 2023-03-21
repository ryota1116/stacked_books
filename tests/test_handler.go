package tests

import (
	"io/ioutil"
	"net/http"
	"testing"
)

type TestHandler struct {
	T *testing.T
}

// PrintErrorFormatFromResponse レスポンスボディの内容を出力する
func (test *TestHandler) PrintErrorFormatFromResponse(response *http.Response) {
	responseBodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	// レスポンス確認
	test.T.Errorf(
		`ステータスコード: %d / ボディ: %s `,
		response.StatusCode,
		responseBodyBytes,
	)
}
