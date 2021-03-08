package persistence

import (
	"../../domain/model"
	"../../domain/repository"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

// Userのインフラ層の構造体
type userPersistence struct{}

// Userデータに関するPersistenceを生成
// 戻り値がinterface型(UserRepository)でなければエラーになる = userPersistence{}をinterface型にした
// インターフェースの中にある同じ名前のメソッドを全て実装すれば、自動的にインターフェイスが実装されたことになる(実装しないとエラーになる)
// 今回で言えば、インターフェイス型UserRepositoryのSignUp, SignIn, ShowUser
func NewUserPersistence() repository.UserRepository {
	// TODO: ここだけ直す！！！
	// https://qiita.com/tono-maron/items/345c433b86f74d314c8d#interface%E3%81%AB%E6%85%A3%E3%82%8C%E3%82%8B
	return &userPersistence{}
}

// 構造体にインターフェイスを実装する書き方
// func (引数 構造体名) 関数名(){
// 	関数の中身
// }
// インターフェイスの実装
func (up userPersistence) SignUp(w http.ResponseWriter, r *http.Request) error {
	db := DbConnect()
	user := model.User{}

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

	return err
}

func (up userPersistence) SignIn(w http.ResponseWriter, r *http.Request) (model.User, error) {
	db := DbConnect()
	user := model.User{}

	json.NewDecoder(r.Body).Decode(&user)

	dbUser := model.User{}
	err := db.Debug().Select([]string{"password"}).Where("email = ?", user.Email).Find(&dbUser).Row().
		Scan(&dbUser.Password) // DBからユーザー取得
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("SQL実行結果")
	fmt.Println(dbUser.Password)

	// TODO: usecaseに移す
	// これが通れば、generateTokenするように分岐させる
	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		fmt.Println("ログインできませんでした")
	} else {
		fmt.Println("ログインできました")
	}
	return user, err
}

//Userを1件取得
func (up userPersistence) ShowUser(w http.ResponseWriter, r *http.Request) model.User {
	db := DbConnect()

	fmt.Println(r.URL) // 「/user/1」とかを取得している

	user := model.User{}
	params := mux.Vars(r) // map[id:1]
	result := db.Debug().First(&user, params["userId"])

	fmt.Println(&result)

	return user
}

// DBサーバーに接続する
func DbConnect() (db *gorm.DB) {
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