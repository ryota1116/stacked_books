package model

// ReadingVolumeCalculator : 読書量を計算するドメインモデル
type ReadingVolumeCalculator struct {
	BookPageCount int `json:"id" validate:"required,max=255"`
}

// BOOK_THICKNESS_PER_PAGE : 本のページ1枚あたりのミリ単位
const BOOK_THICKNESS_PER_PAGE = 1

// calculateInMillimeters : 本の1ページあたりの厚さをミリ単位で計算する
func (readingVolumeCalculator ReadingVolumeCalculator) calculateInMillimeters() int {
	bookPagesNumber := readingVolumeCalculator.ConvertPagesCountIntoPage()

	// ページ1枚 x 1mm で厚さを計算する
	return bookPagesNumber * BOOK_THICKNESS_PER_PAGE
}

// ConvertPagesCountIntoPage : 本の2ページを1枚に変換する
func (readingVolumeCalculator ReadingVolumeCalculator) ConvertPagesCountIntoPage() int {
	return readingVolumeCalculator.BookPageCount / 2
}