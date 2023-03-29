package book

type PublishedYearInterface interface {
	Value() *int
}

type publishedYear struct {
	value *int
}

func NewPublishedYear(value *int) PublishedYearInterface {
	return &publishedYear{value}
}

func (py *publishedYear) Value() *int {
	return py.value
}
