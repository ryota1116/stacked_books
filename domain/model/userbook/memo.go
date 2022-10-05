package userbook

import (
	"fmt"
	"unicode/utf8"
)

// Memo : 本のメモ
type Memo struct {
	Value string
}

// メモの最大文字数
const maxCount = 255

// NewMemo : コンストラクター
func NewMemo(value string) (Memo, error) {
	memoCount := utf8.RuneCountInString(value)

	if memoCount > maxCount {
		return Memo{}, fmt.Errorf("メモは255文字以下で入力ください。")
	}

	return Memo{value}, nil
}
