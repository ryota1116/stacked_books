package model

import (
	"fmt"
	"github.com/go-playground/validator"
	_ "image/png"
	"time"
)

type User struct {
	Id        int64     `json:"id" validate:"required`
	UserName  string    `json:"user_name" validate:"required, max=255"`
	Email     string    `json:"email" validate:"required, unique, max=255, email"`
	Password  string    `json:"password" validate:"required, gt=8, max=255"`
	Avatar    string    `json:"avatar"`
	Role      int       `json:"role" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func main()  {
	user := &User{
		UserName:      "Badger",
		Email:         "Smithgmail.com",
	}
	validate := validator.New()  //インスタンス生成
	errors := validate.Struct(user) //バリデーションを実行し、NGの場合、ここでエラーが返る。

	if errors != nil {
		fmt.Println(errors)
	}
}
