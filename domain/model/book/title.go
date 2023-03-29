package book

type TitleInterface interface {
	Value() string
}

type title struct {
	value string
}

func NewTitle(value string) TitleInterface {
	return &title{value}
}

func (t *title) Value() string {
	return t.value
}
