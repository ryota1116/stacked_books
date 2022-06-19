package user

import "github.com/ryota1116/stacked_books/domain/model"

type SignInDtoGenerator struct {
	User model.User
}

type SignInDto struct {
	Id int
	UserName string
	Email string
	Password string
}

func (sdg SignInDtoGenerator) Execute() SignInDto {
	var signInDto = SignInDto{
		Id:       sdg.User.Id,
		UserName: sdg.User.UserName,
		Email:    sdg.User.Email,
		Password: sdg.User.Password,
	}

	return signInDto
}
