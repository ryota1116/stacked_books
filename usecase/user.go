package usecase

import (
	"../domain/model"
	"../domain/repository"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

const (
	secretKey = "secretKey"
)

// UserにおけるUseCaseのインターフェース
type UserUseCase interface {
	SignUp(user model.User) error
	SignIn(user model.User) (string, error)
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

func (uu userUseCase) SignUp(user model.User) error {
	// bcryptを使ってパスワードをハッシュ化する
	bcryptHashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	dbErr := uu.userRepository.SignUp(user, bcryptHashPassword)
	if dbErr != nil {
		return dbErr
	}
	return dbErr
}

func (uu userUseCase) SignIn(user model.User) (string, error) {
	dbUser, err := uu.userRepository.SignIn(user)

	// TODO: usecaseに移す？
	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		fmt.Println("ログインできませんでした") // レスポンスボディに入れる文字列を返すようにする
		return "", err
	} else {
		fmt.Println("ログインできました")
	}

	token, err := GenerateToken(user)
	return token, err
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
