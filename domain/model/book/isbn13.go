package book

type Isbn13Interface interface {
	Value() *string
}

type isbn13 struct {
	value *string
}

func NewIsbn13(value *string) Isbn13Interface {
	return &isbn13{value}
}

func (s *isbn13) Value() *string {
	return s.value
}
