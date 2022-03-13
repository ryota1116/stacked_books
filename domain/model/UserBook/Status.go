package UserBook

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
func GetBookStatuses() []int {
	return []int{
		WantToRead,
		Reading,
		Done,
	}
}

// Status : ユーザーが登録している本の読書ステータス
type Status struct {
	Value	int	`json:"status"`
}
