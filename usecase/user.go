package usecase

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/ryota1116/stacked_books/domain/model"
	"github.com/ryota1116/stacked_books/domain/repository"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	secretKey = "secretKey"
)

// UserにおけるUseCaseのインターフェース
type UserUseCase interface {
	SignUp(user model.User) (model.User, error)
	SignIn(user model.User) (model.User, error)
	FindOne(userId int) model.User
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

func (uu userUseCase) SignUp(user model.User) (model.User, error) {
	// bcryptを使ってパスワードをハッシュ化する
	bcryptHashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		//return
		fmt.Println(err)
	}

	dbUser, err := uu.userRepository.SignUp(user, bcryptHashPassword)
	return dbUser, err
	//if dbErr != nil {
	//	return dbErr
	//}
	//return dbErr
}


// 「emailで取得したUserのpassword(ハッシュ化されている)」と「クライアントのpassword入力値」を比較する
func (uu userUseCase) SignIn(user model.User) (model.User, error) {
	dbUser, err := uu.userRepository.SignIn(user)

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		fmt.Println("ログインできませんでした") // レスポンスボディに入れる文字列を返すようにする
		return user, err
	} else {
		fmt.Println("ログインできました")
	}

	return dbUser, err
}

func (uu userUseCase) FindOne(userId int) model.User {
	user := uu.userRepository.FindOne(userId)
	return user
}

// GenerateToken : 最後の返り値をerror型(インターフェイス)にすることで、エラーの有無を返す。Goは例外処理が無いため、多値で返すのが基本
// 多値でない(エラーの戻り値が無い)場合、その関数が失敗しないことを期待している？
func GenerateToken(user model.User) (string, error) {
	// 署名生成に使用するアルゴリズムにHS256を使用
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	fmt.Println(token)

	// ペイロードに格納するclaimを作成
	token.Claims = jwt.MapClaims{
		"exp": jwt.TimeFunc().Add(time.Hour * 72).Unix(), // トークンの有効期限
		"iat": jwt.TimeFunc().Unix(), // トークンの生成時間
		"userId": user.Id, // ユーザーID
		"email": user.Email, // メールアドレス
		"password": user.Password, // パスワード
	}
	fmt.Println(token)
	fmt.Println(token.Claims)

	// TODO: シークレットキーを環境変数で持たせる
	// link: https://qiita.com/po3rin/items/740445d21487dfcb5d9f
	// データに対して署名を付与して、文字列にする
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		panic(err.Error())
	}

	return tokenString, nil  // nilでエラーが無かったことを返す
}
