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

func (pm *publishedMonth) Value() *int {
	return pm.value
}
