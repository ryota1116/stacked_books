package user

import (
	"github.com/magiconair/properties/assert"
	modelUser "github.com/ryota1116/stacked_books/domain/model/user"
	_ "net/http/httptest"
	"testing"
	"time"
)

type UserRepositoryMock struct {
}

func (UserRepositoryMock) Create(user modelUser.User) (modelUser.User, error) {
	return modelUser.User{
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

func (UserRepositoryMock) FindOneByEmail(email string) (modelUser.User, error) {
	return modelUser.User{
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

func (UserRepositoryMock) FindOne(userId int) modelUser.User {
	return modelUser.User{
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

	command := UserCreateCommand{
		UserName: "user",
		Email:    "user@example.jp",
		Password: "password",
	}
	// SignUpメソッドの返り値を格納
	user, err := uu.SignUp(command)
	if err != nil {
		t.Fatal(err)
	}

	expected := modelUser.User{
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

	user, err := uu.SignIn("user@example.jp", "password")
	if err != nil {
		t.Errorf("テストに失敗しました。エラーメッセージ: %s", err)
	}

	expected := modelUser.User{
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
