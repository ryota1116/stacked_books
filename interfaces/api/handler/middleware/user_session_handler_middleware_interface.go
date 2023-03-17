package middleware

import (
	"github.com/ryota1116/stacked_books/usecase/user"
	"net/http"
)

type UserSessionHandlerMiddleWareInterface interface {
	CurrentUser(*http.Request) (user.UserDto, error)
}
