package userbook

import (
	"github.com/google/go-cmp/cmp"
	"github.com/ryota1116/stacked_books/domain/model/book"
	"github.com/ryota1116/stacked_books/domain/model/userbook"
	"testing"
	"time"
)

type BookRepositoryMock struct{}

type UserBookRepositoryMock struct{}

func (BookRepositoryMock) FindOneByGoogleBooksId(GoogleBooksId string) book.Book {
	return book.Book{
		Id:             1,
		GoogleBooksId:  "test",
		Title:          "タイトル",
		Description:    "説明文です",
		Image:          "",
		Isbn_10:        "",
		Isbn_13:        "",
		PageCount:      100,
		PublishedYear:  2022,
		PublishedMonth: 8,
		PublishedDate:  10,
		CreatedAt:      time.Date(2022, time.August, 10, 12, 0, 0, 0, time.UTC),
		UpdatedAt:      time.Date(2022, time.August, 10, 12, 0, 0, 0, time.UTC),
	}
}

func (BookRepositoryMock) Save(book2 book.Book) error {
	return nil
}

func (BookRepositoryMock) FindAllByUserId(int) ([]book.Book, error) {
	var books []book.Book
	books = append(books, book.Book{
		Id:             1,
		GoogleBooksId:  "test",
		Title:          "タイトル",
		Description:    "説明文です",
		Image:          "",
		Isbn_10:        "",
		Isbn_13:        "",
		PageCount:      100,
		PublishedYear:  2022,
		PublishedMonth: 8,
		PublishedDate:  10,
		CreatedAt:      time.Date(2022, time.August, 10, 12, 0, 0, 0, time.UTC),
		UpdatedAt:      time.Date(2022, time.August, 10, 12, 0, 0, 0, time.UTC),
	})

	return books, nil
}

func (UserBookRepositoryMock) Save(userbook.UserBook) error {
	return nil
}

// UserBookUseCaseのRegisterUserBookの正常系テスト
func TestUserBookUseCaseRegisterUserBook(t *testing.T) {
	ubu := NewUserBookUseCase(BookRepositoryMock{}, UserBookRepositoryMock{})

	command := UserBookCreateCommand{
		UserId: 1,
		Book: Book{
			GoogleBooksId:  "Wx1dLwEACAAJ",
			Title:          "リーダブルコード",
			Description:    "読んでわかるコードの重要性と方法について解説",
			Isbn10:         "4873115655",
			Isbn13:         "9784873115658",
			PageCount:      237,
			PublishedYear:  2012,
			PublishedMonth: 6,
		},
		UserBook: UserBook{
			Status: 0,
			Memo:   "メモ",
		},
	}

	// userBookUseCaseのRegisterUserBookを実行
	book, userBook, _ := ubu.RegisterUserBook(command)

	// 戻り値である構造体が正しいことをテスト
	if diff := cmp.Diff(book, expectedBook); diff != "" {
		t.Errorf("戻り値の構造体が期待するものではありません。: (-got +want)\n%s", diff)
	}

	if diff := cmp.Diff(userBook, expectedUserBook); diff != "" {
		t.Errorf("戻り値の構造体が期待するものではありません。: (-got +want)\n%s", diff)
	}
}

// expectedBook userBookUseCase.RegisterUserBookの戻り値で期待するBook構造体
// 構造体を定数constに格納することは出来ないので、変数宣言している
var expectedBook = book.Book{
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

// expectedUserBook userBookUseCase.RegisterUserBookの戻り値で期待するUserBook構造体
// 構造体を定数constに格納することは出来ないので、変数宣言している
var expectedUserBook = userbook.UserBook{
	Id:     1,
	UserId: 1,
	BookId: 1,
	Status: 1,
	Memo:   "メモメモメモ",
}
