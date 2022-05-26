package model

// UserBookStatus : 本の読書ステータスのEnumクラス
type UserBookStatus int

const (
	Want UserBookStatus = iota + 1
	Wip
	Done
)

func (status UserBookStatus) GetStatusInt() int {
	switch status {
	case Want:
		return int(Want)
	case Wip:
		return int(Wip)
	case Done:
		return int(Done)
	}
	panic("Unknown value")
}

func (status UserBookStatus) GetStatusString() string {
	switch status {
	case Want:
		return "読みたい"
	case Wip:
		return "読書中"
	case Done:
		return "読了"
	}
	panic("Unknown value")
}