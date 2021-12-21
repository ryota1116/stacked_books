package persistence

import (
	"errors"
	"fmt"
	"github.com/ryota1116/stacked_books/domain/model"
	"github.com/ryota1116/stacked_books/domain/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
func (up userPersistence) SignUp(user model.User, bcryptHashPassword []byte) (model.User, error) {
	db := DbConnect()

	// TODO: playground/validationを使う
	var err error
	if user.UserName == "" {
		err = errors.New("ユーザー名を入力してください")
	}

	user.Password = string(bcryptHashPassword)
	// DBにユーザーを登録
	db.Create(&user)
	fmt.Println(user)
	return user, err
}

func (up userPersistence) SignIn(user model.User) (model.User, error) {
	db := DbConnect()

	dbUser := model.User{}
	err := db.Debug().Select([]string{"password"}).Where("email = ?", user.Email).Find(&dbUser).Row().
		Scan(&dbUser.Password) // DBからユーザー取得
	if err != nil {
		panic(err.Error())
	}

	return dbUser, err
}

//Userを1件取得
func (up userPersistence) ShowUser(params map[string]string) model.User {
	db := DbConnect()

	user := model.User{}
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
