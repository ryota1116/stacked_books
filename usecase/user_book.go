package usecase

import (
	"github.com/ryota1116/stacked_books/domain/model"
	"github.com/ryota1116/stacked_books/domain/repository"
)

type UserBookUseCase interface {
	RegisterUserBook(userBookParameter model.UserBookParameter) model.UserBookParameter
	ReadUserBooks(userId int) []model.Book
	GetUserTotalReadingVolume(userId int) int
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

func (ubu userBookUseCase) RegisterUserBook(userBookParameter model.UserBookParameter) model.UserBookParameter {
	userBookParameter.Book = ubu.bookRepository.FindOrCreateByGoogleBooksId(userBookParameter.GoogleBooksId, userBookParameter)
	userBook := ubu.userBookRepository.CreateOne(userBookParameter)
	return userBook
}

func (ubu userBookUseCase) ReadUserBooks(userId int) []model.Book {
	userBooks := ubu.userBookRepository.ReadUserBooks(userId)
	return userBooks
}

// GetUserTotalReadingVolume : ユーザーの読書量を本の厚さ単位で取得する
func (ubu userBookUseCase) GetUserTotalReadingVolume(userId int) int {
	// ユーザーの読了済みの本を全て取得する
	userBooks := ubu.userBookRepository.ReadUserBooks(userId)

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
