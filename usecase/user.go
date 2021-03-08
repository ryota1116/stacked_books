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

//func VerifyToken(w http.ResponseWriter, r *http.Request) error {
//	// ParseFromRequestで、リクエストヘッダーのAuthorizationからJWTを抽出し、
//	// 抽出したJWTのclaimをparseしてくれる。parseするだけで署名検証とかはしてくれないzv
//	parsedToken, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
//		_, ok := token.Method.(*jwt.SigningMethodHMAC) // 署名アルゴリズムにHS256を使用しているかチェック
//		if !ok {
//			err := errors.New("署名方法が違います")
//			return nil, err
//		}
//		return []byte(secretKey), nil
//	})
//	fmt.Println("parseされたtoken---")
//	fmt.Println(parsedToken)
//
//	if err == nil && parsedToken.Valid {
//		w.WriteHeader(http.StatusOK)
//		json.NewEncoder(w).Encode("認証成功")
//	} else {
//		w.WriteHeader(http.StatusUnauthorized)
//		if err != nil {
//			fmt.Println(err) // key is of invalid type
//		}
//		if !parsedToken.Valid {
//			fmt.Println("トークンが有効ではない")
//		}
//	}
//	return err
//}


// 最後の返り値をerror型(インターフェイス)にすることで、エラーの有無を返す。Goは例外処理が無いため、多値で返すのが基本
// 多値でない(エラーの戻り値が無い)場合、その関数が失敗しないことを期待している？
//func (uu userUseCase) GenerateToken([]*model.User) (string, error) {
//	// 署名生成に使用するアルゴリズムにHS256を使用
//	token := jwt.New(jwt.GetSigningMethod("HS256"))
//	fmt.Println(token)
//
//	user := model.User{}
//
//	// ペイロードに格納するclaimを作成
//	token.Claims = jwt.MapClaims{
//		"exp": jwt.TimeFunc().Add(time.Hour * 72).Unix(), // トークンの有効期限
//		"iat": jwt.TimeFunc().Unix(), // トークンの生成時間
//		"Email": user.Email, // メールアドレス
//		"Password": user.Password, // パスワード
//	}
//	fmt.Println(token)
//	fmt.Println(token.Claims)
//
//	// データに対して署名を付与して、文字列にする
//	tokenString, err := token.SignedString([]byte(secretKey))
//	if err != nil {
//		panic(err.Error())
//	}
//
//	return tokenString, nil  // nilでエラーが無かったことを返す
//}
//
//func (uu userUseCase) VerifyToken(w http.ResponseWriter, r *http.Request) {
//	// ParseFromRequestで、リクエストヘッダーのAuthorizationからJWTを抽出し、
//	// 抽出したJWTのclaimをparseしてくれる。parseするだけで署名検証とかはしてくれないzv
//	parsedToken, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
//		_, ok := token.Method.(*jwt.SigningMethodHMAC) // 署名アルゴリズムにHS256を使用しているかチェック
//		if !ok {
//			err := errors.New("署名方法が違います")
//			return nil, err
//		}
//		return []byte(secretKey), nil
//	})
//	fmt.Println("parseされたtoken---")
//	fmt.Println(parsedToken)
//
//	if err == nil && parsedToken.Valid {
//		w.WriteHeader(http.StatusOK)
//		json.NewEncoder(w).Encode("認証成功")
//	} else {
//		w.WriteHeader(http.StatusUnauthorized)
//		if err != nil {
//			fmt.Println(err) // key is of invalid type
//		}
//		if !parsedToken.Valid {
//			fmt.Println("トークンが有効ではない")
//		}
//	}
//}