package googleBooksApi

import (
	"github.com/magiconair/properties/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	status := m.Run() // テストコードの実行（testing.M.Runで各テストケースが実行され、成功の場合0を返す）。また、各ユニットテストの中でテストデータをinsertすれば良さそう。

	os.Exit(status)   // 0が渡れば成功する。プロセスのkillも実行される。
}

func TestBookHandler_SearchBooks(t *testing.T) {
	t.Run("正常系のテスト", func(t *testing.T) {
		searchWord := "リーダブルコード"


	},
}

// TestSendRequestWithEmptyTitleParameter : リクエストボディのTitleの値が空の場合
func TestSendRequestWithEmptyTitleParameter(t *testing.T) {
	requestParameter := RequestParameter{
		Title: "リーダブルコード",
	}

	searchBooksResponses, _ := Client{}.SendRequest(requestParameter.Title)

	assert.Equal(t, searchBooksResponses, SearchBooksResponses{})
}
