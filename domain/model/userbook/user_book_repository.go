package userbook

type UserBookRepository interface {
	Save(userBook UserBookInterface) error
	FindUserBooksByStatus(userID int, status StatusInterface) ([]UserBookInterface, error)
}
