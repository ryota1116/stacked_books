package userbook

import (
	"fmt"
)

// 本の読書ステータスのEnum
const (
	// WantToRead : 読みたい（= 1）
	WantToRead int = iota + 1
	// Reading : 読書中（= 2）
	Reading
	// Done : 読了（= 3）
	Done
)

// GetBookStatuses : 本の読書ステータス一覧を取得する
func getBookStatuses() []int {
	return []int{
		WantToRead,
		Reading,
		Done,
	}
}

type StatusInterface interface {
	Value() int
}

// Status : 本の読書ステータス
type status struct {
	value int
}

// NewStatus : コンストラクター
func NewStatus(value int) (StatusInterface, error) {
	// TODO: Contain関数を作成する https://zenn.dev/glassonion1/articles/7c7830a269909c
	isValid := func() bool {
		for _, bookStatus := range getBookStatuses() {
			if value == bookStatus {
				return true
			}
		}
		return false
	}

	if !isValid() {
		return &status{}, fmt.Errorf("読書ステータスの値が不正です。 status: %d", value)
	}

	return &status{value}, nil
}

func (s *status) Value() int {
	return s.value
}
