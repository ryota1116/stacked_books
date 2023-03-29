package userbook

import (
	"fmt"
	"unicode/utf8"
)

// メモの最大文字数
const maxCount = 255

type MemoInterface interface {
	Value() *string
	changeMemo(value *string) error
}

// memo : 本のメモ
type memo struct {
	value *string
}

func NewMemo(value *string) (MemoInterface, error) {
	if err := validate(value); err != nil {
		return nil, err
	}

	// &でポインタ型を生成
	return &memo{value}, nil
}

func (m *memo) Value() *string {
	return m.value
}

func (m *memo) changeMemo(value *string) error {
	if err := validate(value); err != nil {
		return err
	}

	m.value = value
	return nil
}

func validate(value *string) error {
	if value != nil {
		memoCount := utf8.RuneCountInString(*value)

		if memoCount > maxCount {
			return fmt.Errorf("メモは255文字以下で入力ください。")
		}
	}

	return nil
}
