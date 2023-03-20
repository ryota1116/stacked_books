package userbook

type UserBookRepository interface {
	Save(userBook UserBookInterface) error
}
