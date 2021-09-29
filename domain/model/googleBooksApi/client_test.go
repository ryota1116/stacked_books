package googleBooksApi

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

// TestSendRequestWithEmptyTitleParameter : リクエストボディのTitleの値が空の場合
func TestSendRequestWithEmptyTitleParameter(t *testing.T) {
	requestParameter := RequestParameter{
		Title: "リーダブルコード",
	}

	searchBooksResponses, _ := Client{}.SendRequest(requestParameter.Title)

	assert.Equal(t, searchBooksResponses, SearchBooksResponses{})
}
