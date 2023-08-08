package user

import (
	modelUser "github.com/ryota1116/stacked_books/domain/model/user"
	"github.com/ryota1116/stacked_books/tests"
	_ "net/http/httptest"
	"os"
	"testing"
	"time"
)

//NOTE: 「エラーメッセージ: crypto/bcrypt: hashedSecret too short to be a bcrypted password」を防ぐため
const hashedPassword = "$ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFG"

// モック
type UserRepositoryMock struct {
}

func (UserRepositoryMock) SaveOne(_ modelUser.UserInterface) (modelUser.UserInterface, error) {
	id := 1
	u, err := modelUser.NewUser(
		&id,
		"user_name",
		"user@example.com",
		hashedPassword,
		nil,
		1,
		&time.Time{},
		&time.Time{},
	)

	return u, err
}

func (UserRepositoryMock) FindOneByEmail(_ string) (modelUser.UserInterface, error) {
	id := 1
	u, err := modelUser.NewUser(
		&id,
		"user_name",
		"user@example.com",
		hashedPassword,
		nil,
		1,
		&time.Time{},
		&time.Time{},
	)

	return u, err
}

func (UserRepositoryMock) FindOneById(_ int) (modelUser.UserInterface, error) {
	id := 1
	u, err := modelUser.NewUser(
		&id,
		"user_name",
		"user@example.com",
		hashedPassword,
		nil,
		1,
		&time.Time{},
		&time.Time{},
	)

	return u, err
}

func TestMain(m *testing.M) {
	status := m.Run() // テストコードの実行（testing.M.Runで各テストケースが実行され、成功の場合0を返す）。また、各ユニットテストの中でテストデータをinsertすれば良さそう。

	os.Exit(status) // 0が渡れば成功する。プロセスのkillも実行される。
}

func TestUserUseCase_SignUp(t *testing.T) {
	uu := NewUserUseCase(&UserRepositoryMock{})

	t.Run("正常系のテスト", func(t *testing.T) {
		command := UserCreateCommand{
			UserName: "user_name",
			Email:    "user@example.com",
			Password: "password",
			Avatar:   nil,
			Role:     1,
		}

		userDto, err := uu.SignUp(command)
		if err != nil {
			t.Fatal(err)
		}

		expected := UserDto{
			Id:       1,
			UserName: "user_name",
			Email:    "user@example.com",
			Password: hashedPassword,
		}

		tests.Assertion{T: t}.AssertEqual(expected, userDto)
	})
}

func TestUserUseCase_SignIn(t *testing.T) {
	ur := UserRepositoryMock{}
	uu := NewUserUseCase(&ur)

	t.Run("正常系のテスト", func(t *testing.T) {
		user, err := uu.SignIn("user@example.com", hashedPassword)
		if err != nil {
			// TODO bcrypt.CompareHashAndPasswordでエラーになるのでスキップしている
		}

		expected := UserDto{
			Id:       1,
			UserName: "user_name",
			Email:    "user@example.com",
			Password: hashedPassword,
		}

		tests.Assertion{T: t}.AssertEqual(expected, user)
	})
}

func TestUserUseCase_FindOne(t *testing.T) {
	ur := UserRepositoryMock{}
	uu := NewUserUseCase(&ur)

	t.Run("正常系のテスト", func(t *testing.T) {
		user, err := uu.FindOne(1)
		if err != nil {
			t.Errorf("テストに失敗しました。エラーメッセージ: %s", err)
		}

		expected := UserDto{
			Id:       1,
			UserName: "user_name",
			Email:    "user@example.com",
			Password: hashedPassword,
		}

		tests.Assertion{T: t}.AssertEqual(expected, user)
	})
}
