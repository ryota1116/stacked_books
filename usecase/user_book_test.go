package usecase

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/ryota1116/stacked_books/domain/model"
	"github.com/ryota1116/stacked_books/domain/model/dto"
	"github.com/ryota1116/stacked_books/tests/expected/api/user_book_use_case"
	"strings"
	"testing"
)

type BookRepositoryMock struct {}

func (BookRepositoryMock) FindOrCreateByGoogleBooksId(dto.RegisterUserBookRequestParameter) model.Book {
	return model.Book{
		GoogleBooksId:  "Wx1dLwEACAAJ",
		Title:          "リーダブルコード",
		Description:    "読んでわかるコードの重要性と方法について解説",
		Isbn_10:        "4873115655",
		Isbn_13:        "9784873115658",
		PageCount:      237,
		PublishedYear:  2012,
		PublishedMonth: 6,
		PublishedDate:  0,
	}
}

type UserBookRepositoryMock struct {}

func (UserBookRepositoryMock) CreateOne(int, int, dto.RegisterUserBookRequestParameter) model.UserBook {
	return model.UserBook{
		UserId:    1,
		BookId:    1,
		Status:    0,
		Memo:      "メモメモメモ",
	}
}

// UserBookUseCaseのRegisterUserBookの正常系テスト
func TestUserBookUseCaseRegisterUserBook(t *testing.T) {
	brm := BookRepositoryMock{}
	ubrm := UserBookRepositoryMock{}
	ubu := NewUserBookUseCase(brm, ubrm)

	// 構造体の中身を検証しているならこの記述が活きてくる気がするが、、
	bodyReader := strings.NewReader(`{
		"google_books_id": "Wx1dLwEACAAJ",
		"title": "リーダブルコード",
		"authors": ["Dustin Boswell","Trevor Foucher"],
		"description": "読んでわかるコードの重要性と方法について解説",
		"isbn_10": "4873115655",
		"isbn_13": "9784873115658",
		"page_count": 237,
		"published_year": 2012,
		"published_month": 6,
	
		"status": 0,
		"memo": "メモメモメモ"
	}`)

	// json文字列をRegisterUserBookRequestParameter構造体に変換
	registerUserBookRequestParams := dto.RegisterUserBookRequestParameter{}
	err := json.NewDecoder(bodyReader).Decode(&registerUserBookRequestParams)
	if err != nil {
		fmt.Println(err)
	}

	expected := user_book_use_case.ExpectedRegisterUserBookResponse

	// userBookUseCaseのRegisterUserBookを実行
	userBookResponse := ubu.RegisterUserBook(1, registerUserBookRequestParams)

	// 戻り値である構造体が正しいことをテスト
	if diff := cmp.Diff(userBookResponse, expected); diff != "" {
		t.Errorf("戻り値の構造体が期待するものではありません。: (-got +want)\n%s", diff)
	}
}
