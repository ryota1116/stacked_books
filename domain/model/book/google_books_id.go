package book

type GoogleBooksIdInterface interface {
	Value() string
}

type googleBooksId struct {
	value string
}

func NewGoogleBooksId(value string) GoogleBooksIdInterface {
	return &googleBooksId{value}
}

func (gbi *googleBooksId) Value() string {
	return gbi.value
}
