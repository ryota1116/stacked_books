package handler

import (
	"encoding/json"
	"github.com/ryota1116/stacked_books/domain/model"
	httpResponse "github.com/ryota1116/stacked_books/handler/http/response"
	ur "github.com/ryota1116/stacked_books/handler/http/response/user"
	"github.com/ryota1116/stacked_books/usecase"
	"io/ioutil"
	"net/http"
)

const (
	secretKey = "secretKey"
)

type UserHandler interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
	ShowUser(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	userUseCase usecase.UserUseCase
}

// Userデータに関するHandlerを生成
// userHandlerをinterface型(UserHandler)にした
func NewUserHandler(uu usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: uu,
	}
}

// uhはuserHandler型の構造体 → つまりUserHandler(インターフェイス型)
func (uh userHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	responseBodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httpResponse.Return500Response(w, err)
		return
	}

	// リクエストをUserの構造体に変換
	user := model.User{}
	if err := json.Unmarshal(responseBodyBytes, &user); err != nil {
		httpResponse.Return500Response(w, err)
		return
	}

	userDto, err := uh.userUseCase.SignUp(user)
	if err != nil {
		httpResponse.Return500Response(w, err)
		return
	}

	token, err := usecase.GenerateToken(userDto)
	if err != nil {
		httpResponse.Return500Response(w, err)
		return
	}

	httpResponse.Response{
		StatusCode: http.StatusOK,
		ResponseBody: ur.SignUpResponseGenerator{
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

	user := model.User{}
	json.NewDecoder(r.Body).Decode(&user)

	userDto, err := uh.userUseCase.SignIn(user)
	if err != nil {
		httpResponse.Return500Response(w, err)
		return
	}

	token, err := usecase.GenerateToken(userDto)
	if err != nil {
		httpResponse.Return500Response(w, err)
		return
	}

	httpResponse.Response{
		StatusCode: http.StatusOK,
		ResponseBody: ur.SignInResponseGenerator{
			UserDto: userDto,
			Token:   token,
		}.Execute(),
	}.ReturnResponse(w)
}

func (uh userHandler) ShowUser(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return
	}

	user = uh.userUseCase.FindOne(user.Id)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
