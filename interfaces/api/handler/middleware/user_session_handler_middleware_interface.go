package middleware

import (
	"github.com/ryota1116/stacked_books/domain/model/user"
	"net/http"
)

type UserSessionHandlerMiddleWareInterface interface {
	CurrentUser(*http.Request) user.User
}
