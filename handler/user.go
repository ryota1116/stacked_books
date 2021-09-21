package handler

import (
	"encoding/json"
	"fmt"
	"github.com/ryota1116/stacked_books/domain/model"
	"github.com/ryota1116/stacked_books/handler/middleware"
	"github.com/ryota1116/stacked_books/usecase"
	"io/ioutil"
	"net/http"
)

const (
	secretKey = "secretKey"
)

type UserHandler interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
	ShowUser(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	userUseCase usecase.UserUseCase
}

// Userデータに関するHandlerを生成
// userHandlerをinterface型(UserHandler)にした
func NewUserHandler(uu usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: uu,
	}
}

// uhはuserHandler型の構造体 → つまりUserHandler(インターフェイス型)
func (uh userHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	responseBodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	// リクエストをUserの構造体に変換
	user := model.User{}
	if err := json.Unmarshal(responseBodyBytes, user); err != nil {
		panic(err)
	}

	// アプリケーションのバリデーションエラーを受け取り、JSONでレスポンスする
	code, errmap := model.UserValidate(user)
	if len(errmap) != 0 {
		errResponse := model.RespondErrJson(code, errmap)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(errResponse)
	} else {
		dbUser, err := uh.userUseCase.SignUp(user)
		if err != nil {
			// TODO: DB側のバリデーションエラーを受け取り、JSONでレスポンスする
			fmt.Println(err)
		}

		// 正常時のレスポンス
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dbUser)
	}
}

func (uh userHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	json.NewDecoder(r.Body).Decode(&user)

	dbUser, err := uh.userUseCase.SignIn(user)
	// tokenを返す
	token, err := usecase.GenerateToken(dbUser)

	if err != nil {
		fmt.Println(err)
	} else {
		middleware.SetUserSession(w, dbUser)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(token) // 生成したトークンをリクエストボディで返してみる
	}
}

func (uh userHandler) ShowUser(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return
	}

	user = uh.userUseCase.FindOne(user.Id)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
