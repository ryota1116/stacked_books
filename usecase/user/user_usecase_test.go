package user

import (
	"github.com/magiconair/properties/assert"
	"github.com/ryota1116/stacked_books/domain/model/user"
	_ "github.com/ryota1116/stacked_books/domain/repository"
	_ "net/http/httptest"
	"testing"
	"time"
)

type UserRepositoryMock struct {
}

func (ur *UserRepositoryMock) SignUp(user user.User, bcryptHashPassword []byte) (user.User, error) {
	return user.User{
		Id:        1,
		UserName:  user.UserName,
		Email:     user.Email,
		Password:  user.Password,
		Avatar:    "",
		Role:      0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}, nil
}

func (ur *UserRepositoryMock) SignIn(user user.User) (user.User, error) {
	return user.User{
		Id:        1,
		UserName:  "user",
		Email:     "user@example.jp",
		Password:  "password",
		Avatar:    "",
		Role:      0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}, nil
}

func (ur *UserRepositoryMock) FindOne(int) user.User {
	return user.User{
		Id:        1,
		UserName:  "user",
		Email:     "user@example.jp",
		Password:  "password",
		Avatar:    "",
		Role:      0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}

func TestUserHandler_SignUp(t *testing.T) {
	ur := UserRepositoryMock{}
	uu := NewUserUseCase(&ur)

	user := user.User{
		UserName: "user_name",
		Email:    "user@example.jp",
		Password: "password",
	}
	// SignUpメソッドの返り値を格納
	user, err := uu.SignUp(user)
	if err != nil {
		t.Fatal(err)
	}

	expected := user.User{
		Id:        1,
		UserName:  user.UserName,
		Email:     user.Email,
		Password:  user.Password,
		Avatar:    "",
		Role:      0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	assert.Equal(
		t,
		expected,
		user,
		"テストに失敗しました。")
}

func TestUserHandler_SignIn(t *testing.T) {
	ur := UserRepositoryMock{}
	uu := NewUserUseCase(&ur)

	user := user.User{
		UserName: "user_name",
		Email:    "user@example.jp",
		Password: "password",
	}

	user, err := uu.SignIn(user)
	if err != nil {
		t.Errorf("テストに失敗しました。エラーメッセージ: %s", err)
	}

	expected := user.User{
		Id:        1,
		UserName:  "user",
		Email:     "user@example.jp",
		Password:  "password",
		Avatar:    "",
		Role:      0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	assert.Equal(
		t,
		expected,
		user,
		"テストに失敗しました。")
}

func TestUserHandler_ShowUser(t *testing.T) {
}
