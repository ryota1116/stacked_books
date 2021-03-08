package repository

import (
	"net/http"
	"../model"
)

// interfaceを定義し、技術的関心事を扱うinfra層がrepositoryの処理を実装する（依存性逆転）
type UserRepository interface {
	// *型名でポイント型になる
	// *型名でUser型へのポイント型
	SignUp(w http.ResponseWriter, r *http.Request) error
	SignIn(w http.ResponseWriter, r *http.Request) (model.User, error)
	ShowUser(w http.ResponseWriter, r *http.Request) model.User
}
