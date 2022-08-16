package user

import "github.com/ryota1116/stacked_books/domain/model"

type UserDtoGenerator struct {
	User model.User
}

type UserDto struct {
	Id       int
	UserName string
	Email    string
	Password string
}

func (sdg UserDtoGenerator) Execute() UserDto {
	var userDto = UserDto{
		Id:       sdg.User.Id,
		UserName: sdg.User.UserName,
		Email:    sdg.User.Email,
		Password: sdg.User.Password,
	}

	return userDto
}
