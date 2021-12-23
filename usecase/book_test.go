package usecase

import (
	"github.com/ryota1116/stacked_books/domain/model/googleBooksApi"
	"os"
	"testing"
)

type googleBooksAPIClientMock struct {
	
}

func (client googleBooksApi.Client) SendRequest(searchWord string) (googleBooksApi.ResponseBodyFromGoogleBooksAPI, error) {
	return googleBooksApi.ResponseBodyFromGoogleBooksAPI{

	}
}

func TestMain(m *testing.M) {
	status := m.Run() // テストコードの実行（testing.M.Runで各テストケースが実行され、成功の場合0を返す）。また、各ユニットテストの中でテストデータをinsertすれば良さそう。

	os.Exit(status)   // 0が渡れば成功する。プロセスのkillも実行される。
}

// インターフェイスを満たしているかテストする？(メソッドが変わった時に)
// => それは型付けで担保できるのでは？
func SearchBooksTest()  {

}

// ユースケース層のSearchBooksメソッドのテスト
func TestBookUseCase_SearchBooks(t *testing.T) {

}
