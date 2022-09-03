package book

import (
	"github.com/magiconair/properties/assert"
	"github.com/ryota1116/stacked_books/domain/model/google-books-api"
	"os"
	"testing"
)

//  IGoogleBooksAPIClientのモック
type googleBooksAPIClientMock struct{}

// IGoogleBooksAPIClient.SendRequestをモックしている
func (client googleBooksAPIClientMock) SendRequest(searchWord string) (google_books_api.ResponseBodyFromGoogleBooksAPI, error) {
	return google_books_api.ResponseBodyFromGoogleBooksAPI{
		Items: []google_books_api.Item{
			google_books_api.Item{
				ID: "Wx1dLwEACAAJ",
				VolumeInfo: google_books_api.VolumeInfo{
					Title: "リーダブルコード",
					Authors: []string{
						"Dustin Boswell",
						"Trevor Foucher",
					},
					PublishedDate: "2012-06",
					Description:   "読んでわかるコードの重要性と方法について解説",
					IndustryIdentifiers: []google_books_api.IndustryIdentifier{
						{
							Type:       "ISBN_10",
							Identifier: "4873115655",
						},
						{
							Type:       "ISBN_13",
							Identifier: "9784873115658",
						},
					},
					PageCount: 237,
				},
			},
		},
	}, nil
}

func TestMain(m *testing.M) {
	status := m.Run() // テストコードの実行（testing.M.Runで各テストケースが実行され、成功の場合0を返す）。また、各ユニットテストの中でテストデータをinsertすれば良さそう。

	os.Exit(status) // 0が渡れば成功する。プロセスのkillも実行される。
}

// インターフェイスを満たしているかテストする？(メソッドが変わった時に)
// => それは型付けで担保できるのでは？
func TestBookUseCase_SearchBooks(t *testing.T) {
	useCase := NewBookUseCase(googleBooksAPIClientMock{})

	expected := google_books_api.ResponseBodyFromGoogleBooksAPI{
		Items: []google_books_api.Item{
			google_books_api.Item{
				ID: "Wx1dLwEACAAJ",
				VolumeInfo: google_books_api.VolumeInfo{
					Title: "リーダブルコード",
					Authors: []string{
						"Dustin Boswell",
						"Trevor Foucher",
					},
					PublishedDate: "2012-06",
					Description:   "読んでわかるコードの重要性と方法について解説",
					IndustryIdentifiers: []google_books_api.IndustryIdentifier{
						{
							Type:       "ISBN_10",
							Identifier: "4873115655",
						},
						{
							Type:       "ISBN_13",
							Identifier: "9784873115658",
						},
					},
					PageCount: 237,
				},
			},
		},
	}

	t.Run("正常系のテスト", func(t *testing.T) {
		responseFromGoogleBooksAPI, err := useCase.SearchBooks("リーダブルコード")
		if err != nil {
			t.Errorf(`テストが失敗しました。エラーメッセージ: %d`, err)
		}

		assert.Equal(
			t,
			responseFromGoogleBooksAPI,
			expected)
	})
}
