package model

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go/request"
	_ "image/png"
	"log"
	"net/http"

	"time"
	//"encoding/json"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/go-openapi/errors"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	secretKey = "secretKey"
)

type User struct {
	Id        int64     `json:"id"`
	UserName  string    `json:"user_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Avatar    string    `json:"avatar"`
	Role      int       `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// 最後の返り値をerror型(インターフェイス)にすることで、エラーの有無を返す。Goは例外処理が無いため、多値で返すのが基本
// 多値でない(エラーの戻り値が無い)場合、その関数が失敗しないことを期待している？
func generateToken(user User) (string, error) {
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

// TODO: 何をクリアしたら認証されたものとするのか整理すること(tokenが存在するか、有効期限内か、署名の検証、useridとpasswordが正しいか)
func verifyToken(w http.ResponseWriter, r *http.Request) {
	// ParseFromRequestで、リクエストヘッダーのAuthorizationからJWTを抽出し、
	// 抽出したJWTのclaimをparseしてくれる。parseするだけで署名検証とかはしてくれないzv
	parsedToken, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC) // 署名アルゴリズムにHS256を使用しているかチェック
		if !ok {
			err := errors.New(401, "署名方法が違います") // APIのエラーを生成
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

func SignUp(w http.ResponseWriter, r *http.Request) {
	db := dbConnect()

	var user User

	fmt.Println("-----リクエストボディ-----")
	fmt.Println(r.Body)

	// リクエストボディをデコードする
	json.NewDecoder(r.Body).Decode(&user)

	// bcryptを使ってパスワードをハッシュ化する
	bcryptHashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("-----リクエストボディをデコードした結果-----")
	fmt.Println(user)

	// DBにユーザーを登録
	if r.Method == "POST" {
		query := "INSERT INTO USERS(user_name, email, password) VALUES(?, ?, ?)"
		// TODO: Execは戻り値がないから、変数に格納できない
		db.Debug().Exec(query, user.UserName, user.Email, bcryptHashPassword)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	db := dbConnect()
	var user User

	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println("リクエストボディをデコードした結果")
	fmt.Println(user)

	var dbUser User
	err := db.Debug().Select([]string{"password"}).Where("email = ?", user.Email).Find(&dbUser).Row().
		Scan(&dbUser.Password)  // DBからユーザー取得
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("SQL実行結果")
	fmt.Println(dbUser.Password)

	// これが通れば、generateTokenするように分岐させる
	//err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		fmt.Println("ログインできませんでした")
	} else {
		fmt.Println("ログインできました")

	}

	// JWTを生成
	tokenString, err := generateToken(user)
	if err != nil {
		panic(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
	json.NewEncoder(w).Encode(tokenString) // 生成したトークンをリクエストボディで返してみる
}


//Userを1件取得
func ShowUser(w http.ResponseWriter, r *http.Request)  {
	db := dbConnect()

	fmt.Println(r.URL) // 「/user/1」とかを取得している

	user := User{}
	params := mux.Vars(r) // map[id:1]
	result := db.Debug().First(&user, params["userId"])

	fmt.Println(&result)
}

// webサーバーに接続する
func StartWebServer() error {
	router := mux.NewRouter().StrictSlash(true)

	// エンドポイント(リクエストを処理して、レスポンスを返す)
	router.HandleFunc("/signup", SignUp).Methods("POST")
	router.HandleFunc("/signin", SignIn).Methods("POST")
	router.HandleFunc("/user/authenticate", verifyToken).Methods("POST")
	router.HandleFunc("/user/{userId:[0-9]+}", ShowUser).Methods("GET")

	log.Println("サーバー起動 : 8080 port で受信")
	return http.ListenAndServe(fmt.Sprintf(":%d", 8080), router)
}

// DBサーバーに接続する
func dbConnect() (db *gorm.DB) {

	// DB接続
	dsn := "root@tcp(127.0.0.1:3306)/stacked_books_development?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	// DBにテーブルが存在するか確認（存在すればtrueを返す）
	dbPresence := db.Migrator().HasTable("users")
	fmt.Println(dbPresence)

	return db
}
