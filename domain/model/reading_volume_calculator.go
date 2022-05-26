package model

// ReadingVolumeCalculator : 読書量を計算するドメインモデル
type ReadingVolumeCalculator struct {
	BookPageCount int `json:"id" validate:"required,max=255"`
}

// BookThicknessPerPage : 本のページ1枚あたりのミリ単位
const BookThicknessPerPage = 1

// CalculateInMillimeters : 本の1ページあたりの厚さをミリ単位で計算する
func (readingVolumeCalculator ReadingVolumeCalculator) CalculateInMillimeters() int {
	bookPagesNumber := readingVolumeCalculator.ConvertPagesCountIntoPage()

	// ページ1枚 x 1mm で厚さを計算する
	return bookPagesNumber * BookThicknessPerPage
}

// ConvertPagesCountIntoPage : 本の2ページを1枚に変換する
func (readingVolumeCalculator ReadingVolumeCalculator) ConvertPagesCountIntoPage() int {
	return readingVolumeCalculator.BookPageCount / 2
}
