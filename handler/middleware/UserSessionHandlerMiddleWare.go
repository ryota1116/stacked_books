package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/ryota1116/stacked_books/domain/model"
	"github.com/ryota1116/stacked_books/infra/persistence"
	"github.com/ryota1116/stacked_books/usecase"
	"net/http"
	"strconv"
	"time"
)

const (
	secretKey = "secretKey"
)

type UserSessionHandlerMiddleWare struct {
	userUseCase usecase.UserUseCase
}

// 認証が通らないとメッセージとリターンを返す（認証失敗時にどのページに繊維するとかはどこで定義する？）
// https://journal.lampetty.net/entry/implementing-middleware-with-http-package-in-go
func VerifyUserToken(w http.ResponseWriter, r *http.Request) {
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
		err := json.NewEncoder(w).Encode("認証成功")
		if err != nil {

		}

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

// CurrentUser : セッションからログイン中のユーザー情報を取得する
func CurrentUser(r *http.Request) model.User {
	// ParseFromRequestでリクエストヘッダーのAuthorizationからJWTを抽出し、抽出したJWTのclaimをparseしてくれる。
	parsedToken, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC) // 署名アルゴリズムにHS256を使用しているかチェック
		if !ok {
			err := errors.New("署名方法が違います")
			return nil, err
		}
		return []byte(secretKey), nil
	})

	userPersistence := persistence.NewUserPersistence()
	userUseCase := usecase.NewUserUseCase(userPersistence)

	//TODO: Goではmodel.Userまたはnilみたいな戻り値は設定できない？
	if err == nil && parsedToken.Valid {
		claims := parsedToken.Claims.(jwt.MapClaims)
		userId := int(claims["userId"].(float64))
		return userUseCase.FindOne(userId)
	} else {
		if err != nil {
			fmt.Println(err) // key is of invalid type
		}
		if !parsedToken.Valid {
			fmt.Println("トークンが有効ではない")
			return userUseCase.FindOne(1)
		}
	}
	return model.User{}
}

func SetUserSession(w http.ResponseWriter, user model.User) {
	expiration := time.Now()
	expiration.AddDate(0, 0, 7)
	cookie := http.Cookie{
		Name:       "user_id",
		Value:      strconv.Itoa(int(user.Id)),
		Expires:    expiration,
	}
	//cookie := http.Cookie{
	//	Name:       "user_session_key",
	//	Value:      uuid.Generate(uuid.Bits),
	//	Expires:    expiration,
	//}

	http.SetCookie(w, &cookie)
}

// いらないかも
func ReadUserSession(w http.ResponseWriter, r *http.Request)  {
}