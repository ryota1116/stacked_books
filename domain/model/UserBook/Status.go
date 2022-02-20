package UserBook

// 本の読書ステータスのEnum
const (
	// WantToRead : 読みたい（= 1）
	WantToRead int = iota + 1
	// READING : 読書中（= 2）
	READING
	// DONE : 読了（= 3）
	DONE
)

// GetBookStatuses : 本の読書ステータス一覧を取得する
func GetBookStatuses() []int {
	return []int{
		WantToRead,
		READING,
		DONE,
	}
}

// Status : ユーザーが登録している本の読書ステータス
type Status struct {
	Value	int	`json:"status"`
}
