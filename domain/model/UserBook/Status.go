package UserBook

// 本の読書ステータスのEnum
const (
	// WANT_TO_READ : 読みたい（= 1）
	WANT_TO_READ int = iota + 1
	// READING : 読書中（= 2）
	READING
	// DONE : 読了（= 3）
	DONE
)

// GetBookStatuses : 本の読書ステータス一覧を取得する
func GetBookStatuses() []int {
	return []int{WANT_TO_READ,READING,DONE}
}

// Status : ユーザーが登録している本の読書ステータス
type Status struct {
	Value	int	`json:"status"`
}
