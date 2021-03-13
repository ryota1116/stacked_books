package handler

import (
	"../domain/model"
	"../usecase"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/mux"
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
	user := model.User{}
	json.NewDecoder(r.Body).Decode(&user)

	// TODO: バリデーションエラーを受け取り、JSONでレスポンスする
	dbUser, err := uh.userUseCase.SignUp(user)
	if err != nil {
		fmt.Println(err)
		// json.NewEncoder(w).Encode(err)
	}

	json.NewEncoder(w).Encode(dbUser)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
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