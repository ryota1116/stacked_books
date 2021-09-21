package repository

import (
	"github.com/ryota1116/stacked_books/domain/model"
)

// interfaceを定義し、技術的関心事を扱うinfra層がrepositoryの処理を実装する（依存性逆転）
type UserRepository interface {
	// *型名でポイント型になる
	// *型名でUser型へのポイント型
	SignUp(user model.User, bcryptHashPassword []byte) (model.User, error)
	SignIn(user model.User) (model.User, error)
	FindOne(userId int) model.User
}
