package userbook

import (
	"github.com/ryota1116/stacked_books/domain/model/book"
	"github.com/ryota1116/stacked_books/domain/model/userbook"
	book2 "github.com/ryota1116/stacked_books/usecase/book"
	"os"
	"reflect"
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
		"test",
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
		"test",
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
		command := UserBookCreateCommand{
			UserId: 1,
			Book: Book{
				GoogleBooksId:  "Wx1dLwEACAAJ",
				Title:          "リーダブルコード",
				Description:    &description,
				Image:          nil,
				Isbn10:         nil,
				Isbn13:         nil,
				PageCount:      237,
				PublishedYear:  &publishedYear,
				PublishedMonth: &publishedMonth,
				PublishedDate:  &publishedDate,
			},
			UserBook: UserBook{
				Status: 0,
				Memo:   &memo,
			},
		}

		bookDto, userBookDto, _ := ubu.RegisterUserBook(command)

		expectedBookDto := book2.BookDto{
			Id:             0,
			GoogleBooksId:  "",
			Title:          "リーダブルコード",
			Description:    &description,
			Image:          nil,
			Isbn10:         nil,
			Isbn13:         nil,
			PageCount:      237,
			PublishedYear:  &publishedYear,
			PublishedMonth: &publishedMonth,
			PublishedDate:  &publishedDate,
		}

		expectedUserBookDto := UserBookCreateCommand{
			UserId: 1,
			Book: Book{
				GoogleBooksId:  "Wx1dLwEACAAJ",
				Title:          "リーダブルコード",
				Description:    &description,
				Image:          nil,
				Isbn10:         nil,
				Isbn13:         nil,
				PageCount:      237,
				PublishedYear:  &publishedYear,
				PublishedMonth: &publishedMonth,
				PublishedDate:  &publishedDate,
			},
			UserBook: UserBook{
				Status: 0,
				Memo:   &memo,
			},
		}

		// 戻り値である構造体が正しいことをテスト
		if reflect.DeepEqual(bookDto, expectedBookDto) {
			t.Errorf(`bookDtoの比較が失敗しました。`)
		}

		if reflect.DeepEqual(userBookDto, expectedUserBookDto) {
			t.Errorf(`userBookDtoの比較が失敗しました。`)
		}

		// TODO: 差分がある場合に、cmpを使って差分を出力できるようにしたい
		//if diff := cmp.Diff(book, expectedBook); diff != "" {
		//	t.Errorf("戻り値の構造体が期待するものではありません。: (-got +want)\n%s", diff)
		//}
	})
}

// UserBookUseCaseのRegisterUserBookの正常系テスト
func TestUserBookUseCase_FindUserBooksByUserId(t *testing.T) {
	ubu := NewUserBookUseCase(BookRepositoryMock{}, UserBookRepositoryMock{})

	t.Run("正常系のテスト", func(t *testing.T) {
		bookDto, _ := ubu.FindUserBooksByUserId(1)

		description := "説明文です"
		publishedYear := 2022
		publishedMonth := 8
		publishedDate := 10
		expectedBookDto := book2.BookDto{
			Id:             0,
			GoogleBooksId:  "",
			Title:          "リーダブルコード",
			Description:    &description,
			Image:          nil,
			Isbn10:         nil,
			Isbn13:         nil,
			PageCount:      237,
			PublishedYear:  &publishedYear,
			PublishedMonth: &publishedMonth,
			PublishedDate:  &publishedDate,
		}

		// 戻り値である構造体が正しいことをテスト
		if reflect.DeepEqual(bookDto, expectedBookDto) {
			t.Errorf(`bookDtoの比較が失敗しました。`)
		}

		// TODO: 差分がある場合に、cmpを使って差分を出力できるようにしたい
		//if diff := cmp.Diff(book, expectedBook); diff != "" {
		//	t.Errorf("戻り値の構造体が期待するものではありません。: (-got +want)\n%s", diff)
		//}
	})
}
