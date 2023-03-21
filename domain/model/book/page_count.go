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

func (pc *pageCount) Value() int {
	return pc.value
}
