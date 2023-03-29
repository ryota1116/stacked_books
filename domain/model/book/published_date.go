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

func (pd *publishedDate) Value() *int {
	return pd.value
}
