package test_assertion

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"testing"
)

// CompareResponseBodyWithJsonFile : レスポンスボディのjson文字列と期待するjson文字列を比較してテストする
func CompareResponseBodyWithJsonFile(t *testing.T, responseBody io.ReadCloser, filePath string) {
	// レスポンスボディを[]byte型に変換
	responseBodyBytes, err := ioutil.ReadAll(responseBody)
	if err != nil {
		panic(err)
	}

	// レスポンスボディのjson文字列をインデントする
	var actual bytes.Buffer
	err = json.Indent(&actual, responseBodyBytes, "", "  ")
	if err != nil {
		t.Fatalf("レスポンスボディのjson文字列のインデントに失敗しました。 '%#v'", err)
	}

	// ファイルの中身を読み込む
	readFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Fatalf("unexpected error while opening file '%#v'", err)
	}

	// JSON文字列の比較
	assert.JSONEq(t, string(readFile), actual.String())
}
