package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/mux"
	"github.com/ryota1116/stacked_books/domain/model"
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
	if err := json.Unmarshal(responseBodyBytes, &user); err != nil {
		panic(err)
	}

	// アプリ側のバリデーションエラーを受け取り、JSONでレスポンスする
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

	token, err := uh.userUseCase.SignIn(user)

	if err != nil {
		fmt.Println(err)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(token) // 生成したトークンをリクエストボディで返してみる
	}
}

func (uh userHandler) ShowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // map[id:1]
	user := uh.userUseCase.ShowUser(params)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func VerifyToken(w http.ResponseWriter, r *http.Request) {
	// ParseFromRequestでリクエストヘッダーのAuthorizationからJWTを抽出し、抽出したJWTのclaimをparseしてくれる。
	parsedToken, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC) // 署名アルゴリズムにHS256を使用しているかチェック
		if !ok {
			err := errors.New("署名方法が違います")
			return nil, err
		}
		return []byte(secretKey), nil
	})
	fmt.Println("parseされたtoken---")
	fmt.Println(parsedToken)

	if err == nil && parsedToken.Valid {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("認証成功")
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		if err != nil {
			fmt.Println(err) // key is of invalid type
		}
		if !parsedToken.Valid {
			fmt.Println("トークンが有効ではない")
		}
	}
}