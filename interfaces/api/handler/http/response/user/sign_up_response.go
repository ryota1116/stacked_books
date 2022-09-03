package user

import "github.com/ryota1116/stacked_books/usecase/user"

type SignUpResponseGenerator struct {
	UserDto user.UserDto
	Token   string
}

type SignUpResponse struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func (sirg SignUpResponseGenerator) Execute() SignUpResponse {
	return SignUpResponse{
		UserName: sirg.UserDto.UserName,
		Email:    sirg.UserDto.Email,
		Token:    sirg.Token,
	}
}
