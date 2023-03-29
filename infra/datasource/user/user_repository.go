package user

import (
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
	Id        *int   `gorm:"primaryKey"`
	UserName  string `validate:"required,max=255"`
	Email     string `validate:"required,max=255,email"`
	Password  string `validate:"required,gte=8,max=255"`
	Avatar    *string
	Role      int
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func (up userPersistence) Save(u userEntity.UserInterface) (userEntity.UserInterface, error) {
	db := datasource.DbConnect()
	uRecord := user{
		UserName: u.UserName().Value(),
		Email:    u.Email().Value(),
		Password: u.Password().Value(),
		Avatar:   u.Avatar().Value(),
		Role:     u.Role().Value(),
	}

	// DBにユーザーを登録
	if err := db.Create(&uRecord).Error; err != nil {
		return nil, err
	}

	u, err := userEntity.NewUser(
		uRecord.Id,
		uRecord.UserName,
		uRecord.Email,
		uRecord.Password,
		uRecord.Avatar,
		uRecord.Role,
		uRecord.CreatedAt,
		uRecord.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (up userPersistence) FindOneByEmail(email string) (userEntity.UserInterface, error) {
	db := datasource.DbConnect()
	uRecord := user{}

	if err := db.Where("email = ?", email).
		First(&uRecord).Error; err != nil {
		return nil, err
	}

	u, err := userEntity.NewUser(
		uRecord.Id,
		uRecord.UserName,
		uRecord.Email,
		uRecord.Password,
		uRecord.Avatar,
		uRecord.Role,
		uRecord.CreatedAt,
		uRecord.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// FindOne Userを1件取得
func (up userPersistence) FindOne(userId int) (userEntity.UserInterface, error) {
	db := datasource.DbConnect()
	uRecord := user{}

	if err := db.Where("id = ?", userId).
		First(&uRecord).Error; err != nil {
		return nil, err
	}

	u, err := userEntity.NewUser(
		uRecord.Id,
		uRecord.UserName,
		uRecord.Email,
		uRecord.Password,
		uRecord.Avatar,
		uRecord.Role,
		uRecord.CreatedAt,
		uRecord.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return u, nil
}
