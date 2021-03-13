package model

import (
	_ "image/png"
	"time"
)

type User struct {
	Id        int64     `json:"id"`
	UserName  string    `json:"user_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Avatar    string    `json:"avatar"`
	Role      int       `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

