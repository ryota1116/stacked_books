package userbook

type UserBookRepository interface {
	Save(userBook UserBook) error
}
