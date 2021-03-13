package handler

import (
	"encoding/json"
	"github.com/magiconair/properties/assert"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"../domain/model"
)

type UserUseCaseMock struct {
}

func (uu *UserUseCaseMock) SignUp(user model.User) (model.User, error) {
	return model.User{
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

func (uu *UserUseCaseMock) SignIn(user model.User) (string, error) {
	return string("token"), nil
}

func (uu *UserUseCaseMock) ShowUser(params map[string]string) model.User {
	return model.User{
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

// TODO: レスポンスボディの型・値、ステータスコードをテスト
func TestUserHandler_SignUp(t *testing.T) {
	uu := UserUseCaseMock{}
	uh := NewUserHandler(&uu)

	// リクエストボディ
	bodyReader := strings.NewReader(`{
		UserName:  "user_name",
		Email:     "user@example.jp",
		Password:  "password",
	}`)

	r := httptest.NewRequest("POST", "/signup", bodyReader)
	w := httptest.NewRecorder()
	uh.SignUp(w, r)

	response := w.Result() //レスポンスを代入

	// ステータスコードのテスト
	if response.StatusCode != 200 {
		t.Errorf(`レスポンスのステータスコードは %d でした`, response.StatusCode)
	}

	// []byte型に変換
	responseBodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// []byte型を構造体に格納
	var user model.User
	if err := json.Unmarshal(responseBodyBytes, &user); err != nil {
		panic(err)
	}

	assert.Equal(t, user.Id, int64(1), `ユーザーIDが`)
}

func TestUserHandler_SignIn(t *testing.T) {
}

func TestUserHandler_ShowUser(t *testing.T) {
}
