package user

import (
	"github.com/ryota1116/stacked_books/tests"
	userUseCase "github.com/ryota1116/stacked_books/usecase/user"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// モックを導入
type UserUseCaseMock struct {
}

func (uu *UserUseCaseMock) SignUp(userUseCase.UserCreateCommand) (userUseCase.UserDto, error) {
	return userUseCase.UserDto{
		Id:       1,
		UserName: "user_name",
		Email:    "user@example.com",
		Password: "password",
	}, nil
}

func (uu *UserUseCaseMock) SignIn(string, string) (userUseCase.UserDto, error) {
	return userUseCase.UserDto{
		Id:       1,
		UserName: "user_name",
		Email:    "user@example.com",
		Password: "password",
	}, nil
}

func (uu *UserUseCaseMock) FindOne(int) (userUseCase.UserDto, error) {
	return userUseCase.UserDto{
		Id:       1,
		UserName: "user_name",
		Email:    "user@example.com",
		Password: "password",
	}, nil
}

func (uu *UserUseCaseMock) GenerateToken(userUseCase.UserDto) (string, error) {
	return "sample_token", nil
}

// テストで期待するレスポンスボディJSON文字列のファイルパス
const expectedJsonDirectory = "/tests/expected/api/userHandler"

func TestMain(m *testing.M) {
	// テストコードの実行（testing.M.Runで各テストケースが実行され、成功の場合0を返す）
	// => また各ユニットテストの中でテストデータをinsertすれば良さそう。
	status := m.Run()

	// 0が渡れば成功する。プロセスのkillも実行される。
	os.Exit(status)
}

// TODO: レスポンスボディの型・値、ステータスコードをテスト
func TestUserHandler_SignUp(t *testing.T) {
	uh := NewUserHandler(&UserUseCaseMock{}) // モックを注入

	// jsonファイルの絶対パスを取得(TODO: ローカル用の取得になっているので修正する)
	_, testFilePath, _, _ := runtime.Caller(0)
	projectRootDir := filepath.Join(filepath.Dir(testFilePath), "..", "..", "..", "..", "..", "..")

	t.Run("正常系のテスト", func(t *testing.T) {
		testHandler := tests.TestHandler{T: t}

		// リクエスト
		body := strings.NewReader(`{
			"user_name": "user_name",
			"email": "user@example.com",
			"password": "password"
		}`)
		r := httptest.NewRequest("POST", "/signup", body)
		w := httptest.NewRecorder()
		uh.SignUp(w, r)
		response := w.Result() // レスポンスを代入

		// ステータスコードのテスト
		if response.StatusCode != 200 {
			testHandler.PrintErrorFormatFromResponse(response)
		}

		expectedJsonFilePath := filepath.Join(
			projectRootDir,
			expectedJsonDirectory+"/signup/200_response.json",
		)

		// レスポンスボディのjson文字列をテスト
		testHandler.CompareResponseBodyWithJsonFile(
			response.Body,
			expectedJsonFilePath,
		)
	})
}

func TestUserHandler_SignIn(t *testing.T) {
	uh := NewUserHandler(&UserUseCaseMock{}) // モックを注入

	// jsonファイルの絶対パスを取得(TODO: ローカル用の取得になっているので修正する)
	_, testFilePath, _, _ := runtime.Caller(0)
	projectRootDir := filepath.Join(filepath.Dir(testFilePath), "..", "..", "..", "..", "..", "..")

	t.Run("正常系のテスト", func(t *testing.T) {
		testHandler := tests.TestHandler{T: t}

		// リクエスト
		body := strings.NewReader(`{
			"email": "user@example.com",
			"password": "password"
		}`)
		r := httptest.NewRequest("POST", "/signin", body)
		w := httptest.NewRecorder()
		uh.SignIn(w, r)
		response := w.Result() // レスポンスを代入

		// ステータスコードのテスト
		if response.StatusCode != 200 {
			testHandler.PrintErrorFormatFromResponse(response)
		}

		expectedJsonFilePath := filepath.Join(
			projectRootDir,
			expectedJsonDirectory+"/signin/200_response.json",
		)

		// レスポンスボディのjson文字列をテスト
		testHandler.CompareResponseBodyWithJsonFile(
			response.Body,
			expectedJsonFilePath,
		)
	})
}
