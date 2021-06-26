package model

import (
	"gorm.io/gorm"
	_ "image/png"
	"time"
)

// TODO: アプリケーション側でunique制約を付けるには？（DBにアクセスする必要が出てくる）
type User struct {
	gorm.Model
	Id        int64     `json:"id" gorm:"primaryKey"`
	UserName  string    `json:"user_name" validate:"required,max=255"`
	Email     string    `json:"email" validate:"required,max=255,email"`
	Password  string    `json:"password" validate:"required,gte=8,max=255"`
	Avatar    string    `json:"avatar"`
	Role      int       `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Books     []Book    `gorm:"many2many:user_books;"`
}
