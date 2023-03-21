package user

import (
	"fmt"
	userEntity "github.com/ryota1116/stacked_books/domain/model/user"
	"github.com/ryota1116/stacked_books/infra/datasource"
	"time"
)

// Userのインフラ層の構造体
type userPersistence struct{}

// NewUserPersistence Userデータに関するPersistenceを生成
// 戻り値がinterface型(UserRepository)でなければエラーになる = userPersistence{}をinterface型にした
// インターフェースの中にある同じ名前のメソッドを全て実装すれば、自動的にインターフェイスが実装されたことになる(実装しないとエラーになる)
// 今回で言えば、インターフェイス型UserRepositoryのSignUp, SignIn, ShowUser
func NewUserPersistence() userEntity.UserRepository {
	// TODO: ここだけ直す！！！
	// https://qiita.com/tono-maron/items/345c433b86f74d314c8d#interface%E3%81%AB%E6%85%A3%E3%82%8C%E3%82%8B
	return &userPersistence{}
}

type user struct {
	Id        int
	UserName  string
	Email     string
	Password  string
	Avatar    string
	Role      int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// Create 構造体にインターフェイスを実装する書き方
// func (引数 構造体名) 関数名(){
// 	関数の中身
// }
// インターフェイスの実装
func (up userPersistence) Create(user userEntity.User) error {
	db := datasource.DbConnect()

	// TODO: Domainに移す(DBの制約は必須だが最終防衛手段みたいなものなので
	//       業務知識としてDomainに書くのがいい)
	//var err error
	//if user.UserName == "" {
	//	err = errors.New("ユーザー名を入力してください")
	//}

	// DBにユーザーを登録
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (up userPersistence) FindOneByEmail(email string) (userEntity.User, error) {
	db := datasource.DbConnect()
	user := user{}

	err := db.Where("email = ?", email).First(&user).Error
	fmt.Println("========")
	fmt.Println(user)
	if err != nil {
		return userEntity.User{}, nil
	}

	//if err := db.Where("email = ?", email).
	//	First(&user).Error; err != nil {
	//	return userEntity.User{}, err
	//}
	// err := db.Debug().Select([]string{"password"}).Where("email = ?", user.Email).Find(&dbUser).Row().Scan(&dbUser.Password) // DBからユーザー取得

	return userEntity.User{
		Id:        user.Id,
		UserName:  user.UserName,
		Email:     user.Email,
		Password:  user.Password,
		Avatar:    user.Avatar,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
		Books:     nil,
	}, nil
}

// FindOne Userを1件取得
func (up userPersistence) FindOne(userId int) (userEntity.User, error) {
	db := datasource.DbConnect()
	user := user{}

	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		return userEntity.User{}, err
	}

	return userEntity.User{
		Id:        user.Id,
		UserName:  user.UserName,
		Email:     user.Email,
		Password:  user.Password,
		Avatar:    user.Avatar,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
		Books:     nil,
	}, nil
}
