package tests

import (
	"github.com/ryota1116/stacked_books/domain/model/user"
	"github.com/ryota1116/stacked_books/interfaces/api/handler/middleware"
	"net/http"
)

type userSessionHandlerMiddleWareTest struct{}

func NewUserSessionHandlerMiddleWareTest() middleware.UserSessionHandlerMiddleWareInterface {
	return userSessionHandlerMiddleWareTest{}
}

func (userSessionHandlerMiddleWareTest) CurrentUser(r *http.Request) user.User {
	return user.User{Id: 1}
}
