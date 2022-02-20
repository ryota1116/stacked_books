package tests

import (
	"github.com/ryota1116/stacked_books/domain/model"
	"github.com/ryota1116/stacked_books/handler/middleware"
	"net/http"
)

type userSessionHandlerMiddleWareTest struct {}

func NewUserSessionHandlerMiddleWareTest() middleware.UserSessionHandlerMiddleWareInterface {
	return userSessionHandlerMiddleWareTest{}
}

func (userSessionHandlerMiddleWareTest) CurrentUser(r *http.Request) model.User {
	return model.User{Id: 1}
}
