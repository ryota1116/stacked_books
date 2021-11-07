package persistence

import (
	"errors"
	"fmt"
	"github.com/ryota1116/stacked_books/domain/model"
	"github.com/ryota1116/stacked_books/domain/repository"
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
	// emailでUserを取得
	err := db.Where("email = ?", user.Email).First(&dbUser).Error // DBからユーザー取得
	// err := db.Debug().Select([]string{"password"}).Where("email = ?", user.Email).Find(&dbUser).Row().Scan(&dbUser.Password) // DBからユーザー取得

	if err != nil {
		panic(err.Error())
	}

	return dbUser, err
}

//Userを1件取得
func (up userPersistence) FindOne(userId int) model.User {
	db := DbConnect()

	user := model.User{}
	result := db.Debug().First(&user, userId)

	fmt.Println(&result)

	return user
}


