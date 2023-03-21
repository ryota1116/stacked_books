package book

type PublishedMonthInterface interface {
	Value() *int
}

type publishedMonth struct {
	value *int
}

func NewPublishedMonth(value *int) PublishedMonthInterface {
	return &publishedMonth{value}
}

func (s *publishedMonth) Value() *int {
	return s.value
}
