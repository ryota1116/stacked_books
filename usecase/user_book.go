package usecase

import (
	"github.com/ryota1116/stacked_books/domain/model"
	"github.com/ryota1116/stacked_books/domain/model/dto"
	"github.com/ryota1116/stacked_books/domain/repository"
)

type UserBookUseCase interface {
	RegisterUserBook(int, dto.RegisterUserBookRequestParameter) dto.RegisterUserBookResponse
	FindUserBooksByUserId(userId int) ([]model.Book, error)
	GetUserTotalReadingVolume(w http.ResponseWriter, r *http.Request)
}

type userBookUseCase struct {
	bookRepository repository.BookRepository
	userBookRepository repository.UserBookRepository
}

func NewUserBookUseCase(br repository.BookRepository, ubr repository.UserBookRepository) UserBookUseCase {
	return &userBookUseCase{
		bookRepository:     br,
		userBookRepository: ubr,
	}
}

// RegisterUserBook : UserBooksレコードを作成する
func (uub userBookUseCase) RegisterUserBook(userId int, registerUserBookRequestParameter dto.RegisterUserBookRequestParameter) dto.RegisterUserBookResponse {
	// GoogleBooksIDからBookレコードを検索し、存在しなければ作成する
	book := uub.bookRepository.FindOrCreateByGoogleBooksId(registerUserBookRequestParameter)
	// UserBooksレコードを作成する
	userBook := uub.userBookRepository.CreateOne(userId, book.Id, registerUserBookRequestParameter)
	// RegisterUserBookResponse構造体を生成する
	userBookResponse := dto.BuildRegisterUserBookResponse(book, userBook)

	return userBookResponse
}

// FindUserBooksByUserId : ログイン中のユーザーが登録している本の一覧を取得する
func (ubu userBookUseCase) FindUserBooksByUserId(userId int) ([]model.Book, error) {
	userBooks, err := ubu.userBookRepository.FindAllByUserId(userId)
	return userBooks, err
}

// GetUserTotalReadingVolume : ユーザーの読書量を本の厚さ単位で取得する
func (ubu userBookUseCase) GetUserTotalReadingVolume(userId int) int {
	// ユーザーの読了済みの本を全て取得する
	readingStatus := model.UserBookStatus(model.Done).GetStatusInt()
	userBooks := ubu.userBookRepository.FindUserBooksWithReadingStatus(userId, readingStatus)

	// ユーザーの読書量を計算する
	var userTotalReadingVolume int
	for _, userBook := range userBooks {
		readingVolumeCalculator := model.ReadingVolumeCalculator{
			BookPageCount: userBook.PageCount,
		}
		// 本の1ページあたりの厚さをミリ単位で計算する
		userTotalReadingVolume += readingVolumeCalculator.CalculateInMillimeters()
	}

	return userTotalReadingVolume
}
