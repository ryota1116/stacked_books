package model

import "time"

// TODO: ドメインモデルをORMのEntityの用に使ってしまっているから、 persistence/userbook/entity.go?を作成する。
type UserBook struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	UserId    int       `json:"user_id"`
	BookId    int       `json:"book_id"`
	Status    int       `json:"status"`
	Memo      string    `json:"memo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Book Book `json:"userbook"`
}
