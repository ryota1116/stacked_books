package handler

import (
	"encoding/json"
	"github.com/magiconair/properties/assert"
	"io/ioutil"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/ryota1116/stacked_books/domain/model"
)

// モックを導入
type UserUseCaseMock struct {
}
// モック型でプロダクションコードの
func (uu *UserUseCaseMock) SignUp(user model.User) (model.User, error) {
	return model.User{
		Id:        1,
		UserName:  "user_name",
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
		UserName:  "user_name",
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
	// uh := NewUserHandler(&UserUseCaseMock{}) と同義
	uu := UserUseCaseMock{} // モック生成
	uh := NewUserHandler(&uu)

	// リクエストボディ
	bodyReader := strings.NewReader(`{
		"user_name": "user_name",
		"email": "user@example.jp",
		"password": "password",
	}`)

	// TODO: 書き換えてもエラーにならない
	r := httptest.NewRequest("POST", "/signup", bodyReader)
	w := httptest.NewRecorder()
	uh.SignUp(w, r)

	response := w.Result() // uh.SignUp(w, r)は戻り値無いし、レスポンスを代入してテストする

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

	assert.Equal(t, user.Id, int64(1), `ユーザーIDが正しいこと`)
	assert.Equal(t, user.UserName, "user_name", `ユーザー名が正しいこと`)
	assert.Equal(t, user.Password, "password", `ユーザー名が正しいこと`)
}

func TestUserHandler_SignIn(t *testing.T) {
	TestUserHandler_SignUp(t)

	uu := UserUseCaseMock{}
	uh := NewUserHandler(&uu)

	// SignInのリクエストボディ
	signInBodyReader := strings.NewReader(`{
		"email": "user@example.jp",
		"password": "password",
	}`)
	r := httptest.NewRequest("POST", "/signin", signInBodyReader)
	w := httptest.NewRecorder()
	uh.SignIn(w, r)

	response := w.Result() //レスポンスを代入

	// ステータスコードのテスト
	if response.StatusCode != 200 {
		t.Errorf(`レスポンスのステータスコードは %d でした`, response.StatusCode)
	}

	//assert.Equal(t, &response.Body, "string", `ユーザーIDが正しいこと`)
}

func TestUserHandler_ShowUser(t *testing.T) {
}


