package config

import (
	"os"
)

//func setUp(t *testing.T, envs map[string]string)  {
//	for key, value := range envs {
//
//	}
//
//	if err := os.LookupEnv(k) {
//
//	}
//}

// SetLocalDBConfig : ローカルDBの環境変数を設定する
func SetLocalDBConfig()  {
	os.Setenv("DB_USER", "user")
	os.Setenv("DB_PASSWORD", "password")
	os.Setenv("DB_NAME", "stacked_books_development")
}

// SetTestDBConfig : テストDBの環境変数を設定する
func SetTestDBConfig()  {
	os.Setenv("DB_USER", "user")
	os.Setenv("DB_PASSWORD", "password")
	os.Setenv("DB_NAME", "stacked_books_test")
}

// ReadConfig : 環境変数を読み取る
func ReadConfig(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok  {

	} else {
		return val
	}
	return val
}
