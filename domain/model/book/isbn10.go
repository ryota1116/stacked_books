package book

type Isbn10Interface interface {
	Value() *string
}

type isbn10 struct {
	value *string
}

func NewIsbn10(value *string) Isbn10Interface {
	return &isbn10{value}
}

func (i *isbn10) Value() *string {
	return i.value
}
