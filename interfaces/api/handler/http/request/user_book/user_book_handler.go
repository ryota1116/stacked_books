package user_book

import (
	"encoding/json"
	RegisterUserBooks2 "github.com/ryota1116/stacked_books/interfaces/api/handler/http/request/user_book/register_user_books"
	httpResponse "github.com/ryota1116/stacked_books/interfaces/api/handler/http/response"
	"github.com/ryota1116/stacked_books/interfaces/api/handler/http/response/user_book"
	"github.com/ryota1116/stacked_books/interfaces/api/handler/http/response/user_book/find_user_books"
	middleware2 "github.com/ryota1116/stacked_books/interfaces/api/handler/middleware"
	userBookUseCase "github.com/ryota1116/stacked_books/usecase/userbook"
	"net/http"
)

type UserBookHandler interface {
	RegisterUserBook(w http.ResponseWriter, r *http.Request)
	FindUserBooks(w http.ResponseWriter, r *http.Request)
}

type userBookHandler struct {
	userBookUseCase              userBookUseCase.UserBookUseCase
	userSessionHandlerMiddleWare middleware2.UserSessionHandlerMiddleWareInterface
}

func NewUserBookHandler(
	ubu userBookUseCase.UserBookUseCase,
	ushmw middleware2.UserSessionHandlerMiddleWareInterface) UserBookHandler {
	return &userBookHandler{
		userBookUseCase:              ubu,
		userSessionHandlerMiddleWare: ushmw,
	}
}

// RegisterUserBook : booksを参照→同じのあればそれを使って、user_booksを作成
func (ubh userBookHandler) RegisterUserBook(w http.ResponseWriter, r *http.Request) {
	// JSONのリクエストボディを構造体に変換する
	requestBody := RegisterUserBooks2.RequestBody{}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		httpResponse.Response{
			StatusCode:   http.StatusInternalServerError,
			ResponseBody: err.Error(),
		}.ReturnResponse(w)
		return
	}

	// リクエストボディ構造体のバリデーションを実行
	isValid, errMsg := RegisterUserBooks2.FormValidator{
		RequestBody: requestBody,
	}.Validate()
	if !isValid {
		httpResponse.Response{
			StatusCode:   http.StatusUnprocessableEntity,
			ResponseBody: errMsg,
		}.ReturnResponse(w)
		return
	}

	// ログイン中のユーザーを取得する
	currentUser := ubh.userSessionHandlerMiddleWare.CurrentUser(r)

	command := userBookUseCase.UserBookCreateCommand{
		UserId: currentUser.Id,
		Book: userBookUseCase.Book{
			GoogleBooksId:  requestBody.Book.GoogleBooksId,
			Title:          requestBody.Book.Title,
			Description:    requestBody.Book.Description,
			Isbn10:         requestBody.Book.Isbn10,
			Isbn13:         requestBody.Book.Isbn13,
			PageCount:      requestBody.Book.PageCount,
			PublishedYear:  requestBody.Book.PublishedYear,
			PublishedMonth: requestBody.Book.PublishedMonth,
			PublishedDate:  requestBody.Book.PublishedDate,
		},
		UserBook: userBookUseCase.UserBook{
			Status: requestBody.UserBook.Status,
			Memo:   requestBody.UserBook.Memo,
		},
	}

	// UserBooksレコードを作成する
	book, userBook := ubh.userBookUseCase.RegisterUserBook(command)

	httpResponse.Response{
		StatusCode: http.StatusOK,
		ResponseBody: user_book.RegisterUserBookResponseGenerator{
			Book:     book,
			UserBook: userBook,
		}.Execute(),
	}.ReturnResponse(w)
}

// FindUserBooks : ログイン中のユーザーが登録している本の一覧を取得する
func (ubh userBookHandler) FindUserBooks(w http.ResponseWriter, r *http.Request) {
	// セッション情報からUserを取得
	ushm := middleware2.NewUserSessionHandlerMiddleWare()
	user := ushm.CurrentUser(r)

	userBooksDto, err := ubh.userBookUseCase.FindUserBooksByUserId(user.Id)
	if err != nil {
		httpResponse.Return500Response(w, err)
		return
	}

	httpResponse.Response{
		StatusCode: http.StatusOK,
		ResponseBody: find_user_books.FindUserBooksResponseGenerator{
			UserBooksDto: userBooksDto,
		}.Execute(),
	}.ReturnResponse(w)
}
