package usecase

import (
	//"net/http/httptest"
	"testing"
	//"../domain/repository"
	"../domain/model"
)

type UserRepositoryMock struct {
}

func (ur *UserRepositoryMock) SignUp(user model.User, bcryptHashPassword []byte) (model.User, error) {
	return model.User{
		UserName: "user",
		Email: "user@example.jp",
		Password: "password",
	}, nil
}

func TestUserHandler_SignUp(t *testing.T) {
	ur := UserRepositoryMock{}
	uu := NewUserUseCase(ur)
	dbUser, err := uu.SignUp()
}

func TestUserHandler_SignIn(t *testing.T) {
}

func TestUserHandler_ShowUser(t *testing.T) {
}