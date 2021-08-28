package model

// UserBookStatus : 本の読書ステータスのEnumクラス
type UserBookStatus int

const (
	want UserBookStatus = iota + 1
	wip
	done
)
