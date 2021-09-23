package middleware

import (
	"github.com/ryota1116/stacked_books/domain/model"
	"net/http"
)

type UserSessionHandlerMiddleWareInterface interface {
	CurrentUser(*http.Request) model.User
}
