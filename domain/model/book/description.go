package book

type DescriptionInterface interface {
	Value() *string
}

type description struct {
	value *string
}

func NewDescription(value *string) DescriptionInterface {
	return &description{value}
}

func (d *description) Value() *string {
	return d.value
}
