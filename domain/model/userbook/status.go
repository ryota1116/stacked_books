package userbook

import (
	"fmt"
)

// Status : 本の読書ステータス
type Status struct {
	Value int
}

// 本の読書ステータスのEnum
const (
	// WantToRead : 読みたい（= 1）
	WantToRead int = iota + 1
	// READING : 読書中（= 2）
	READING
	// DONE : 読了（= 3）
	DONE
)

// NewStatus : コンストラクター
func NewStatus(value int) (Status, error) {
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
		return Status{}, fmt.Errorf("読書ステータスの値が不正です。 status: %d", value)
	}

	return Status{value}, nil
}

// GetBookStatuses : 本の読書ステータス一覧を取得する
func getBookStatuses() []int {
	return []int{
		WantToRead,
		READING,
		DONE,
	}
}
