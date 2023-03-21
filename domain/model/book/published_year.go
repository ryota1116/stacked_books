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

func (s *publishedYear) Value() *int {
	return s.value
}
