package user

import (
	"errors"
	"fmt"
	"github.com/ryota1116/stacked_books/domain/model/user"
	"github.com/ryota1116/stacked_books/infra/persistence"
)

// Userのインフラ層の構造体
type userPersistence struct{}

// NewUserPersistence Userデータに関するPersistenceを生成
// 戻り値がinterface型(UserRepository)でなければエラーになる = userPersistence{}をinterface型にした
// インターフェースの中にある同じ名前のメソッドを全て実装すれば、自動的にインターフェイスが実装されたことになる(実装しないとエラーになる)
// 今回で言えば、インターフェイス型UserRepositoryのSignUp, SignIn, ShowUser
func NewUserPersistence() user.UserRepository {
	// TODO: ここだけ直す！！！
	// https://qiita.com/tono-maron/items/345c433b86f74d314c8d#interface%E3%81%AB%E6%85%A3%E3%82%8C%E3%82%8B
	return &userPersistence{}
}

// Create 構造体にインターフェイスを実装する書き方
// func (引数 構造体名) 関数名(){
// 	関数の中身
// }
// インターフェイスの実装
func (up userPersistence) Create(user user.User) (user.User, error) {
	db := persistence.DbConnect()

	// TODO: playground/validationを使う
	var err error
	if user.UserName == "" {
		err = errors.New("ユーザー名を入力してください")
	}

	// DBにユーザーを登録
	db.Create(&user)

	return user, err
}

func (up userPersistence) FindOneByEmail(email string) (user.User, error) {
	db := persistence.DbConnect()

	u := user.User{}
	// emailでUserを取得
	err := db.Debug().Where("email = ?", email).First(&u).Error // DBからユーザー取得
	// err := db.Debug().Select([]string{"password"}).Where("email = ?", user.Email).Find(&dbUser).Row().Scan(&dbUser.Password) // DBからユーザー取得

	if err != nil {
		panic(err.Error())
	}

	return u, err
}

// FindOne Userを1件取得
func (up userPersistence) FindOne(userId int) user.User {
	db := persistence.DbConnect()

	u := user.User{}
	result := db.First(&u, userId)

	fmt.Println(&result)

	return u
}
