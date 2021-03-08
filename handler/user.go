package handler

import (
	"../usecase"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"net/http"
	"errors"
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
	err := uh.userUseCase.SignUp(w, r)
	if err != nil {
		fmt.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (uh userHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	user, err := uh.userUseCase.SignIn(w, r)
	if err != nil {
		fmt.Println(err)
	} else {
		tokenString, err := usecase.GenerateToken(user) // token生成
		if err != nil {
			fmt.Println(err)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		//json.NewEncoder(w).Encode(user)
		json.NewEncoder(w).Encode(tokenString) // 生成したトークンをリクエストボディで返してみる
	}
}

func (uh userHandler) ShowUser(w http.ResponseWriter, r *http.Request) {
	user := uh.userUseCase.ShowUser(w, r) // usecaseを呼んでいるだけ

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

//func (uh userHandler) VerifyToken(w http.ResponseWriter, r *http.Request) {
//
//	err := usecase.VerifyToken
//	if err != nil {
//		w.Header()
//	}
//}

func VerifyToken(w http.ResponseWriter, r *http.Request) {
	// ParseFromRequestで、リクエストヘッダーのAuthorizationからJWTを抽出し、
	// 抽出したJWTのclaimをparseしてくれる。parseするだけで署名検証とかはしてくれないzv
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