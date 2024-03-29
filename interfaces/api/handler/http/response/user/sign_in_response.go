package user

import "github.com/ryota1116/stacked_books/usecase/user"

type SignInResponseGenerator struct {
	UserDto user.UserDto
	Token   string
}

type SignInResponse struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func (sirg SignInResponseGenerator) Execute() SignInResponse {
	return SignInResponse{
		UserName: sirg.UserDto.UserName,
		Email:    sirg.UserDto.Email,
		Token:    sirg.Token,
	}
}
