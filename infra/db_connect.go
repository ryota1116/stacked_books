package infra

import (
	"fmt"
	config "github.com/ryota1116/stacked_books"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Db : 接続先DB（グローバル変数を定義している）
var Db *gorm.DB

type configDB struct {
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// DbConnect : DBサーバーに接続する
func DbConnect() {
	configDB := configDB{
		Host:     "",
		Name:     config.ReadConfig("DB_NAME"),
		User:     config.ReadConfig("DB_USER"),
		Password: config.ReadConfig("DB_PASSWORD"),
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configDB.User,
		configDB.Password,
		configDB.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	Db = db

	if err != nil {
		panic(err.Error())
	}
}

func TestMigration()  {
	config.SetTestDBConfig()

	DbConnect()
}

// BeginTransaction : トランザクションの開始
func BeginTransaction() *gorm.DB {
	config.SetTestDBConfig()

	DbConnect()

	return Db.Begin()
}
