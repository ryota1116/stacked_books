package userbook

import (
	"github.com/ryota1116/stacked_books/domain/model/book"
	"github.com/ryota1116/stacked_books/domain/model/userbook"
	"github.com/ryota1116/stacked_books/tests"
	book2 "github.com/ryota1116/stacked_books/usecase/book"
	"os"
	"testing"
	"time"
)

type BookRepositoryMock struct{}

func (BookRepositoryMock) FindAllByUserId(_ int) ([]book.BookInterface, error) {
	var books []book.BookInterface

	id := 1
	description := "説明文です"
	publishedYear := 2022
	publishedMonth := 8
	publishedDate := 10
	createdAt := time.Date(2022, time.August, 10, 12, 0, 0, 0, time.UTC)
	b, _ := book.NewBook(
		&id,
		"test_id",
		"タイトル",
		&description,
		nil,
		nil,
		nil,
		100,
		&publishedYear,
		&publishedMonth,
		&publishedDate,
		&createdAt,
	)
	books = append(books, b)

	return books, nil
}

func (BookRepositoryMock) FindOneByGoogleBooksId(_ string) (book.BookInterface, error) {
	id := 1
	description := "説明文です"
	publishedYear := 2022
	publishedMonth := 8
	publishedDate := 10
	createdAt := time.Date(2022, time.August, 10, 12, 0, 0, 0, time.UTC)
	b, _ := book.NewBook(
		&id,
		"test_id",
		"タイトル",
		&description,
		nil,
		nil,
		nil,
		100,
		&publishedYear,
		&publishedMonth,
		&publishedDate,
		&createdAt,
	)

	return b, nil
}

func (BookRepositoryMock) Save(_ book.BookInterface) error {
	return nil
}

type UserBookRepositoryMock struct{}

func (UserBookRepositoryMock) Save(_ userbook.UserBookInterface) error {
	return nil
}

func TestMain(m *testing.M) {
	status := m.Run() // テストコードの実行（testing.M.Runで各テストケースが実行され、成功の場合0を返す）。また、各ユニットテストの中でテストデータをinsertすれば良さそう。

	os.Exit(status) // 0が渡れば成功する。プロセスのkillも実行される。
}

// UserBookUseCaseのRegisterUserBookの正常系テスト
func TestUserBookUseCase_RegisterUserBook(t *testing.T) {
	ubu := NewUserBookUseCase(BookRepositoryMock{}, UserBookRepositoryMock{})

	t.Run("正常系のテスト", func(t *testing.T) {
		description := "説明文です"
		publishedYear := 2022
		publishedMonth := 8
		publishedDate := 10
		memo := "メモ"

		expectedBookDto := book2.BookDto{
			Id:             1,
			GoogleBooksId:  "test_id",
			Title:          "タイトル",
			Description:    &description,
			Image:          nil,
			Isbn10:         nil,
			Isbn13:         nil,
			PageCount:      100,
			PublishedYear:  &publishedYear,
			PublishedMonth: &publishedMonth,
			PublishedDate:  &publishedDate,
		}
		expectedUserBookDto := UserBookDto{
			UserId: 1,
			BookId: 1,
			Status: 1,
			Memo:   &memo,
		}

		command := UserBookCreateCommand{
			UserId: 1,
			Book: Book{
				GoogleBooksId:  "test_id",
				Title:          "タイトル",
				Description:    &description,
				Image:          nil,
				Isbn10:         nil,
				Isbn13:         nil,
				PageCount:      100,
				PublishedYear:  &publishedYear,
				PublishedMonth: &publishedMonth,
				PublishedDate:  &publishedDate,
			},
			UserBook: UserBook{
				Status: 1,
				Memo:   &memo,
			},
		}
		bookDto, userBookDto, _ := ubu.RegisterUserBook(command)

		// 戻り値である構造体が正しいことをテスト
		tests.Assertion{T: t}.AssertEqual(expectedBookDto, bookDto)
		tests.Assertion{T: t}.AssertEqual(expectedUserBookDto, userBookDto)
	})
}

// UserBookUseCaseのRegisterUserBookの正常系テスト
func TestUserBookUseCase_FindUserBooksByUserId(t *testing.T) {
	ubu := NewUserBookUseCase(BookRepositoryMock{}, UserBookRepositoryMock{})

	t.Run("正常系のテスト", func(t *testing.T) {
		// 期待値を作成
		var expected []book2.BookDto
		description := "説明文です"
		publishedYear := 2022
		publishedMonth := 8
		publishedDate := 10
		expected = append(expected, book2.BookDto{
			Id:             1,
			GoogleBooksId:  "test_id",
			Title:          "タイトル",
			Description:    &description,
			Image:          nil,
			Isbn10:         nil,
			Isbn13:         nil,
			PageCount:      100,
			PublishedYear:  &publishedYear,
			PublishedMonth: &publishedMonth,
			PublishedDate:  &publishedDate,
		})

		bookDto, _ := ubu.FindUserBooksByUserId(1)

		// 戻り値である構造体が正しいことをテスト
		tests.Assertion{T: t}.AssertEqual(expected, bookDto)
	})
}
