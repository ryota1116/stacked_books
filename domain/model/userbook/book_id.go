package userbook

type BookIdInterface interface {
	Value() int
}

type bookId struct {
	value int
}

func NewBookId(value int) BookIdInterface {
	return &bookId{value}
}

func (s *bookId) Value() int {
	return s.value
}
