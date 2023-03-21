package book

type PublishedDateInterface interface {
	Value() *int
}

type publishedDate struct {
	value *int
}

func NewPublishedDate(value *int) PublishedDateInterface {
	return &publishedDate{value}
}

func (s *publishedDate) Value() *int {
	return s.value
}
