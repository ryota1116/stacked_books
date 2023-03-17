package datasource

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// TODO: この関数をどこから呼ぶべきか？
// DBサーバーに接続する
func DbConnect() (db *gorm.DB) {
	// DB接続
	dsn := "root@tcp(127.0.0.1:3306)/stacked_books_development?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	// DBにテーブルが存在するか確認（存在すればtrueを返す）
	dbPresence := db.Migrator().HasTable("users")
	fmt.Println(dbPresence)

	return db
}
