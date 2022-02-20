package usecase

import (
	"github.com/magiconair/properties/assert"
	"github.com/ryota1116/stacked_books/domain/model/googleBooksApi"
	"github.com/ryota1116/stacked_books/handler/http/request/book/search_books"
	"os"
	"testing"
)

//  IGoogleBooksAPIClientのモック
type googleBooksAPIClientMock struct {}

// IGoogleBooksAPIClient.SendRequestをモックしている
func (client googleBooksAPIClientMock) SendRequest(searchWord string) (googleBooksApi.ResponseBodyFromGoogleBooksAPI, error) {
	return googleBooksApi.ResponseBodyFromGoogleBooksAPI{
		Items:  []googleBooksApi.Item{
			googleBooksApi.Item{
				ID:         "Wx1dLwEACAAJ",
				VolumeInfo: googleBooksApi.VolumeInfo{
					Title:               "リーダブルコード",
					Authors:             []string{
						"Dustin Boswell",
						"Trevor Foucher",
					},
					PublishedDate:       "2012-06",
					Description:         "読んでわかるコードの重要性と方法について解説",
					IndustryIdentifiers: []googleBooksApi.IndustryIdentifier{
						{
							Type:       "ISBN_10",
							Identifier: "4873115655",
						},
						{
							Type:       "ISBN_13",
							Identifier: "9784873115658",
						},
					},
					PageCount:           237,
				},
			},
		},
	}, nil
}

func TestMain(m *testing.M) {
	status := m.Run() // テストコードの実行（testing.M.Runで各テストケースが実行され、成功の場合0を返す）。また、各ユニットテストの中でテストデータをinsertすれば良さそう。

	os.Exit(status)   // 0が渡れば成功する。プロセスのkillも実行される。
}

// インターフェイスを満たしているかテストする？(メソッドが変わった時に)
// => それは型付けで担保できるのでは？
func TestBookUseCase_SearchBooks(t *testing.T) {
	useCase := NewBookUseCase(googleBooksAPIClientMock{})

	expected := googleBooksApi.ResponseBodyFromGoogleBooksAPI{
		Items:  []googleBooksApi.Item{
			googleBooksApi.Item{
				ID:         "Wx1dLwEACAAJ",
				VolumeInfo: googleBooksApi.VolumeInfo{
					Title:               "リーダブルコード",
					Authors:             []string{
						"Dustin Boswell",
						"Trevor Foucher",
					},
					PublishedDate:       "2012-06",
					Description:         "読んでわかるコードの重要性と方法について解説",
					IndustryIdentifiers: []googleBooksApi.IndustryIdentifier{
						{
							Type:       "ISBN_10",
							Identifier: "4873115655",
						},
						{
							Type:       "ISBN_13",
							Identifier: "9784873115658",
						},
					},
					PageCount:           237,
				},
			},
		},
	}

	t.Run("正常系のテスト", func(t *testing.T) {
		requestBody := search_books.RequestBody{Title: "リーダブルコード"}

		responseFromGoogleBooksAPI, err := useCase.SearchBooks(requestBody)
		if err != nil {
			t.Errorf(`テストが失敗しました。エラーメッセージ: %d`, err)
		}

		assert.Equal(
			t,
			responseFromGoogleBooksAPI,
			expected)
	})
}
