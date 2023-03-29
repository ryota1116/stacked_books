package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	user2 "github.com/ryota1116/stacked_books/infra/datasource/user"
	"github.com/ryota1116/stacked_books/usecase/user"
	"net/http"
	"strconv"
	"time"
)

const (
	secretKey = "secretKey"
)

type UserSessionHandlerMiddleWareInterface interface {
	CurrentUser(*http.Request) (user.UserDto, error)
}

type userSessionHandlerMiddleWare struct{}

// TODO: ミドルウェアとしながらもリクエスト処理の前後に挟んでいないので、
//       ミドルウェアとすべきものすべきでないものを分けた上で改修すること。
func NewUserSessionHandlerMiddleWare() UserSessionHandlerMiddleWareInterface {
	return userSessionHandlerMiddleWare{}
}

// 消していいかも
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
func (userSessionHandlerMiddleWare) CurrentUser(r *http.Request) (user.UserDto, error) {
	// ParseFromRequestでリクエストヘッダーのAuthorizationからJWTを抽出し、抽出したJWTのclaimをparseしてくれる。
	parsedToken, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC) // 署名アルゴリズムにHS256を使用しているかチェック
		if !ok {
			err := errors.New("署名方法が違います")
			return nil, err
		}
		return []byte(secretKey), nil
	})

	userPersistence := user2.NewUserPersistence()
	userUseCase := user.NewUserUseCase(userPersistence)

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

	// この処理に入る = エラーということ(つまりerrをnilで返してはいけない)
	return user.UserDto{}, err
}

func SetUserSession(w http.ResponseWriter, user user.UserDto) {
	expiration := time.Now()
	expiration.AddDate(0, 0, 7) // 7日間有効 // TODO: 短くする
	cookie := http.Cookie{
		Name:     "user_id",
		Value:    strconv.Itoa(int(user.Id)),
		Expires:  expiration,
		HttpOnly: true,
	}
	//cookie := http.Cookie{
	//	Name:       "user_session_key",
	//	Value:      uuid.Generate(uuid.Bits),
	//	Expires:    expiration,
	//}

	http.SetCookie(w, &cookie)
}
