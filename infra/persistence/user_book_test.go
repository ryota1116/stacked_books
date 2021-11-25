package persistence

import (
	"github.com/magiconair/properties/assert"
	"github.com/ryota1116/stacked_books/db/testdata"
	"github.com/ryota1116/stacked_books/domain/model"
	"github.com/ryota1116/stacked_books/domain/model/dto"
	"github.com/ryota1116/stacked_books/infra"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	shutdown := setupDB()
	defer shutdown()  // 関数がreturnされた後に遅延実行される（代わりにCleanUpでもいけそう？）

	status := m.Run() // テストコードの実行（testing.M.Runで各テストケースが実行され、成功の場合0を返す）。また、各ユニットテストの中でテストデータをinsertすれば良さそう。

	os.Exit(status)   // 0が渡れば成功する。プロセスのkillも実行される。
}

func setupDB() func() {
	return func() {
		// shutdown
	}
}

func TestUserBookPersistence_CreateOne(t *testing.T) {
	t.Run("正常系のテスト", func(t *testing.T) {
		infra.Db = infra.BeginTransaction()

		// テストデータ投入
		testdata.Exec(infra.Db,"../../db/testdata/user_book_test/insert_user.sql")
		testdata.Exec(infra.Db,"../../db/testdata/user_book_test/insert_book.sql")

		// 必要最低限のパラメーターオブジェクトを用意する
		registerUserBookRequestParameter := dto.RegisterUserBookRequestParameter{
			UserBook: dto.UserBook{
				Status:    0,
				Memo:      "メモ",
			},
		}

		// テスト対象のメソッドを実行
		userBookPersistence := NewUserBookPersistence()
		userBook, _ := userBookPersistence.CreateOne(1, 1, registerUserBookRequestParameter)

		// 作成されたデータを取得する
		dbResult := model.UserBook{}
		infra.Db.Where(
			"user_id = ? AND book_id = ?",
			userBook.UserId, userBook.BookId).First(&dbResult)

		// 期待する結果かが得られたか確認
		assert.Equal(t, dbResult.UserId, 1)
		assert.Equal(t, dbResult.BookId, 1)
		assert.Equal(t, dbResult.Memo, "メモ")

		infra.Db.Rollback()
	})

	t.Run("異常系のテスト", func(t *testing.T) {
		infra.Db = infra.BeginTransaction()

		// テストデータ投入
		testdata.Exec(infra.Db,"../../db/testdata/user_book_test/insert_user.sql")
		testdata.Exec(infra.Db,"../../db/testdata/user_book_test/insert_book.sql")

		// 必要最低限のパラメーターオブジェクトを用意する
		registerUserBookRequestParameter := dto.RegisterUserBookRequestParameter{
			UserBook: dto.UserBook{
				Status:    0,
				Memo:      "メモ",
			},
		}

		// テスト対象のメソッドを実行
		// Book.idに存在しないレコードを指定する
		userBookPersistence := NewUserBookPersistence()
		_, err := userBookPersistence.CreateOne(1, 2, registerUserBookRequestParameter)

		// エラーがnilでないことをテストしている
		if err == nil {
			t.Error("実際の値: ", err, "期待値: ", nil)
		}

		infra.Db.Rollback()
	})
}
