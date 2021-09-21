package usecase

import (
	"github.com/ryota1116/stacked_books/domain/model"
	"github.com/ryota1116/stacked_books/domain/repository"
)

type UserBookUseCase interface {
	RegisterUserBook(userId int, userBookParameter model.UserBookParameter) model.UserBookParameter
  ReadUserBooks(userId int) model.Book
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

func (uub userBookUseCase) RegisterUserBook(userId int, userBookParameter model.UserBookParameter) model.UserBookParameter {
	userBookParameter.Book = uub.bookRepository.FindOrCreateByGoogleBooksId(userBookParameter.GoogleBooksId, userBookParameter)
	userBook := uub.userBookRepository.CreateOne(userId, userBookParameter)

	return userBook
}

func (ubu userBookUseCase) ReadUserBooks(userId int) []model.Book {
	userBooks := ubu.userBookRepository.ReadUserBooks(userId)
	return userBooks
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
