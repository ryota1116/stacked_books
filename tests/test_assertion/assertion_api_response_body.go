package test_assertion

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

// CompareResponseBodyWithJsonFile : レスポンスボディのjson文字列と期待するjson文字列を比較してテストする
func CompareResponseBodyWithJsonFile(t *testing.T, responseBodyBytes []byte, filePath string)  {
	// レスポンスボディのjson文字列をインデントする
	var actual bytes.Buffer
	err := json.Indent(&actual, responseBodyBytes, "", "  ")
	if err != nil {
		t.Fatalf("レスポンスボディのjson文字列のインデントに失敗しました。 '%#v'", err)
	}

	// ファイルの中身を読み込む
	byte, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Fatalf("unexpected error while opening file '%#v'", err)
	}

	// JSON文字列の比較
	assert.JSONEq(t, string(byte), actual.String())
}
