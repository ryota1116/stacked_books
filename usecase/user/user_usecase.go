package user

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/ryota1116/stacked_books/domain/model/user"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	secretKey = "secretKey"
)

// UserUseCase UserにおけるUseCaseのインターフェース
type UserUseCase interface {
	SignUp(command UserCreateCommand) (UserDto, error)
	SignIn(email string, password string) (UserDto, error)
	FindOne(userId int) (UserDto, error)
}

// TODO: 依存する方向てきな？
type userUseCase struct {
	userRepository user.UserRepository
}

// NewUserUseCase Userデータに関するUseCaseを生成
// 戻り値にInterface型を指定
//
func NewUserUseCase(ur user.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: ur,
	}
}

func (uu userUseCase) SignUp(command UserCreateCommand) (UserDto, error) {
	// bcryptを使ってパスワードをハッシュ化する
	bcryptHashPassword, err := bcrypt.GenerateFromPassword([]byte(command.Password), bcrypt.DefaultCost)
	if err != nil {
		//return
		fmt.Println(err)
	}

	u := user.User{
		UserName: command.UserName,
		Email:    command.Email,
		Password: string(bcryptHashPassword),
	}

	err = uu.userRepository.Create(u)

	userDto := UserDtoGenerator{
		User: u,
	}.Execute()

	return userDto, err
}

// SignIn 「emailで取得したUserのpassword(ハッシュ化されている)」と「クライアントのpassword入力値」を比較する
func (uu userUseCase) SignIn(email string, password string) (UserDto, error) {
	user, err := uu.userRepository.FindOneByEmail(email)

	userDto := UserDtoGenerator{
		User: user,
	}.Execute()

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		fmt.Println("ログインできませんでした") // レスポンスボディに入れる文字列を返すようにする
		return userDto, err
	} else {
		fmt.Println("ログインできました")
	}

	return userDto, err
}

func (uu userUseCase) FindOne(userId int) (UserDto, error) {
	user, err := uu.userRepository.FindOne(userId)
	if err != nil {
		return UserDto{}, err
	}

	return UserDtoGenerator{
		User: user,
	}.Execute(), nil
}

// GenerateToken : 最後の返り値をerror型(インターフェイス)にすることで、エラーの有無を返す。Goは例外処理が無いため、多値で返すのが基本
// 多値でない(エラーの戻り値が無い)場合、その関数が失敗しないことを期待している？
func GenerateToken(user UserDto) (string, error) {
	// 署名生成に使用するアルゴリズムにHS256を使用
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	fmt.Println(token)

	// ペイロードに格納するclaimを作成
	token.Claims = jwt.MapClaims{
		"exp":      jwt.TimeFunc().Add(time.Hour * 72).Unix(), // トークンの有効期限
		"iat":      jwt.TimeFunc().Unix(),                     // トークンの生成時間
		"userId":   user.Id,                                   // ユーザーID
		"userName": user.UserName,                             // ユーザー名
		"email":    user.Email,                                // メールアドレス
		"password": user.Password,                             // パスワード
	}
	fmt.Println(token)

	// TODO: シークレットキーを環境変数で持たせる
	// link: https://qiita.com/po3rin/items/740445d21487dfcb5d9f
	// データに対して署名を付与して、文字列にする
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		panic(err.Error())
	}

	return tokenString, nil // nilでエラーが無かったことを返す
}
