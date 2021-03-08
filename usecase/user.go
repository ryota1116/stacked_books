package usecase

import (
	"../domain/model"
	"../domain/repository"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

const (
	secretKey = "secretKey"
)

// UserにおけるUseCaseのインターフェース
type UserUseCase interface {
	SignUp(w http.ResponseWriter, r *http.Request) error
	SignIn(w http.ResponseWriter, r *http.Request) (model.User, error)
	ShowUser(w http.ResponseWriter, r *http.Request) model.User
	//GenerateToken([]*model.User) (string, error)
	//VerifyToken(w http.ResponseWriter, r *http.Request)
}

// TODO: 依存する方向てきな？
type userUseCase struct {
	userRepository repository.UserRepository
}

// Userデータに関するUseCaseを生成
// 戻り値にInterface型を指定
//
func NewUserUseCase(ur repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: ur,
	}
}

func (uu userUseCase) SignUp(w http.ResponseWriter, r *http.Request) error {
	err := uu.userRepository.SignUp(w, r)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (uu userUseCase) SignIn(w http.ResponseWriter, r *http.Request) (model.User, error) {
	user, err := uu.userRepository.SignIn(w, r)
	if err != nil {
		fmt.Println(err)
	}
	return user, err
}

func (uu userUseCase) ShowUser(w http.ResponseWriter, r *http.Request) model.User {
	user := uu.userRepository.ShowUser(w, r)
	return user
}

// 最後の返り値をerror型(インターフェイス)にすることで、エラーの有無を返す。Goは例外処理が無いため、多値で返すのが基本
// 多値でない(エラーの戻り値が無い)場合、その関数が失敗しないことを期待している？
func GenerateToken(user model.User) (string, error) {
	// 署名生成に使用するアルゴリズムにHS256を使用
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	fmt.Println(token)

	// ペイロードに格納するclaimを作成
	token.Claims = jwt.MapClaims{
		"exp": jwt.TimeFunc().Add(time.Hour * 72).Unix(), // トークンの有効期限
		"iat": jwt.TimeFunc().Unix(), // トークンの生成時間
		"Email": user.Email, // メールアドレス
		"Password": user.Password, // パスワード
	}
	fmt.Println(token)
	fmt.Println(token.Claims)

	// データに対して署名を付与して、文字列にする
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		panic(err.Error())
	}

	return tokenString, nil  // nilでエラーが無かったことを返す
}
