package searchUserBooksByResponse

import (
	"github.com/ryota1116/stacked_books/usecase/userbook"
	"time"
)

type ResponseGenerator struct {
	UserBooksDto []userbook.UserBookDto
}

// ResponseBody : 読書ステータスでユーザーが登録している本一覧を取得するAPIのレスポンスボディ構造体
type ResponseBody []searchUserBooksByResponse

type searchUserBooksByResponse struct {
	UserId      int       `json:"user_id"`
	BookId      int       `json:"book_id"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	Image       *string   `json:"image"`
	Status      int       `json:"status"`
	Memo        string    `json:"memo"`
	CreatedAt   time.Time `json:"created_at"`
}

func (rg ResponseGenerator) Execute() ResponseBody {
	// DTOに変換
	var response []searchUserBooksByResponse
	for _, userBookDto := range rg.UserBooksDto {
		r := searchUserBooksByResponse{
			UserId:      userBookDto.UserId,
			BookId:      userBookDto.BookId,
			Title:       userBookDto.BookDto.Title,
			Description: userBookDto.BookDto.Description,
			Image:       nil,
			Status:      userBookDto.Status,
			Memo:        *userBookDto.Memo,
			CreatedAt:   userBookDto.CreatedAt,
		}

		response = append(response, r)
	}
	return response
}
