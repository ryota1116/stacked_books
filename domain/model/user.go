package model

import (
	_ "image/png"
	"time"
)

type User struct {
	Id        int64     `json:"id"`
	UserName  string    `json:"user_name" validate:"required,max=255"`
	Email     string    `json:"email" validate:"required,max=255,email"`
	Password  string    `json:"password" validate:"required,gt=8,max=255"`
	Avatar    string    `json:"avatar"`
	Role      int       `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

