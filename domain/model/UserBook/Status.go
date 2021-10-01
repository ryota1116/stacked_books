package UserBook

const (
	// WANT_TO_READ : 読みたい
	WANT_TO_READ int = iota + 1
	// READING : 読書中
	READING
	// DONE : 読了
	DONE
)

// Status : ユーザーが登録している本の読書ステータス
type Status struct {
	Value	int	`json:"status"`
}
