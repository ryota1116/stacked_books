package userbook

type UserBookRepository interface {
	CreateOne(userBook UserBook) UserBook
}
