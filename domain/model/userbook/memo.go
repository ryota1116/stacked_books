package userbook

import (
	"fmt"
	"unicode/utf8"
)

// メモの最大文字数
const maxCount = 255

type MemoInterface interface {
	Value() *string
}

// memo : 本のメモ
type memo struct {
	value *string
}

// NewMemo : コンストラクター
func NewMemo(value *string) (MemoInterface, error) {
	if value != nil {
		memoCount := utf8.RuneCountInString(*value)

		if memoCount > maxCount {
			return &memo{}, fmt.Errorf("メモは255文字以下で入力ください。")
		}
	}

	// &でポインタ型を生成
	return &memo{value}, nil
}

func (m *memo) Value() *string {
	return m.value
}
