package user

import (
	"github.com/ryota1116/stacked_books/domain/model/user"
)

type UserDtoGenerator struct {
	User user.UserInterface
}

type UserDto struct {
	Id       int
	UserName string
	Email    string
	Password string
}

func (sdg UserDtoGenerator) Execute() UserDto {
	return UserDto{
		Id:       *sdg.User.Id().Value(),
		UserName: sdg.User.UserName().Value(),
		Email:    sdg.User.Email().Value(),
		Password: sdg.User.Password().Value(),
	}
}
