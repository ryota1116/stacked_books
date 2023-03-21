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

func (s *title) Value() string {
	return s.value
}
