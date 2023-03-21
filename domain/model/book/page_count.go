package book

type PageCountInterface interface {
	Value() int
}

type pageCount struct {
	value int
}

func NewPageCount(value int) PageCountInterface {
	return &pageCount{value}
}

func (s *pageCount) Value() int {
	return s.value
}
