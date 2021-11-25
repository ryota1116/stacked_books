package testdata

import (
	"fmt"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
)

// Exec : テストデータの投入
func Exec(db *gorm.DB, queryFile string) {
	// クエリの読み込み
	query := fileOpen(queryFile)
	// SQLの実行
	db.Exec(query)
}

// fileOpen : テスト用のSQLファイルを読み込む
func fileOpen(queryFile string) string {
	f, err := os.Open(queryFile)
	if err != nil{
		fmt.Println("エラー")
	}

	defer f.Close()

	// ファイルの中身を全て読み取る
	byte, err := ioutil.ReadAll(f)
	if err != nil{
		fmt.Println("エラー")
	}

	// 文字列型に変換して返す
	return string(byte)
}
