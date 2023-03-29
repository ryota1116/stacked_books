package user

import (
	"encoding/json"
	"github.com/ryota1116/stacked_books/interfaces/api/handler/http/request/user/sign_in"
	"github.com/ryota1116/stacked_books/interfaces/api/handler/http/request/user/sign_up"
	httpResponse "github.com/ryota1116/stacked_books/interfaces/api/handler/http/response"
	user2 "github.com/ryota1116/stacked_books/interfaces/api/handler/http/response/user"
	userUseCase "github.com/ryota1116/stacked_books/usecase/user"
	"io/ioutil"
	"net/http"
)

const (
	secretKey = "secretKey"
)

type UserHandler interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	userUseCase userUseCase.UserUseCase
}

// NewUserHandler Userデータに関するHandlerを生成
// userHandlerをinterface型(UserHandler)にした
func NewUserHandler(uu userUseCase.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: uu,
	}
}

// SignUp uhはuserHandler型の構造体 → つまりUserHandler(インターフェイス型)
func (uh userHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	responseBodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httpResponse.Return500Response(w, err)
		return
	}

	requestBody := sign_up.RequestBody{}
	if err := json.Unmarshal(responseBodyBytes, &requestBody); err != nil {
		httpResponse.Return500Response(w, err)
		return
	}

	command := userUseCase.UserCreateCommand{
		UserName: requestBody.UserName,
		Email:    requestBody.Email,
		Password: requestBody.Password,
		Avatar:   requestBody.Avatar,
		Role:     requestBody.Role,
	}

	userDto, err := uh.userUseCase.SignUp(command)
	if err != nil {
		httpResponse.Return500Response(w, err)
		return
	}

	token, err := uh.userUseCase.GenerateToken(userDto)
	if err != nil {
		httpResponse.Return500Response(w, err)
		return
	}

	httpResponse.Response{
		StatusCode: http.StatusOK,
		ResponseBody: user2.SignUpResponseGenerator{
			UserDto: userDto,
			Token:   token,
		}.Execute(),
	}.ReturnResponse(w)
}

func (uh userHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Requested-With, Origin, X-Csrftoken, Accept, Cookie")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3002")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	requestBody := sign_in.RequestBody{}
	json.NewDecoder(r.Body).Decode(&requestBody)

	userDto, err := uh.userUseCase.SignIn(requestBody.Email, requestBody.Password)
	if err != nil {
		httpResponse.Return500Response(w, err)
		return
	}

	token, err := uh.userUseCase.GenerateToken(userDto)
	if err != nil {
		httpResponse.Return500Response(w, err)
		return
	}

	httpResponse.Response{
		StatusCode: http.StatusOK,
		ResponseBody: user2.SignInResponseGenerator{
			UserDto: userDto,
			Token:   token,
		}.Execute(),
	}.ReturnResponse(w)
}
