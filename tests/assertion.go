package tests

import (
	"reflect"
	"testing"
)

type Assertion struct {
	T *testing.T
}

// AssertEqual : 構造体などを比較する
func (assertion Assertion) AssertEqual(
	expected interface{},
	actual interface{},
) {
	if !reflect.DeepEqual(expected, actual) {
		assertion.T.Errorf(
			"比較に失敗しました。\n期待値: %+v \n実際の値: %+v",
			expected,
			actual,
		)
	}
}
