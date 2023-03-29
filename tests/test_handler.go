package tests

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
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

// CompareResponseBodyWithJsonFile : レスポンスボディのjson文字列と期待するjson文字列を比較してテストする
func (test *TestHandler) CompareResponseBodyWithJsonFile(responseBody io.ReadCloser, filePath string) {
	// レスポンスボディを[]byte型に変換
	responseBodyBytes, err := ioutil.ReadAll(responseBody)
	if err != nil {
		panic(err)
	}

	// レスポンスボディのjson文字列をインデントする
	var actual bytes.Buffer
	err = json.Indent(&actual, responseBodyBytes, "", "  ")
	if err != nil {
		test.T.Fatalf("レスポンスボディのjson文字列のインデントに失敗しました。 '%#v'", err)
	}

	// ファイルの中身を読み込む
	readFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		test.T.Fatalf("unexpected error while opening file '%#v'", err)
	}

	// JSON文字列の比較
	assert.JSONEq(test.T, string(readFile), actual.String())
}
