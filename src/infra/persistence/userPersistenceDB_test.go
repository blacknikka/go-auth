package persistence

import (
	"errors"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestDB(t *testing.T) {
	t.Run("CreateUser正常系", func(t *testing.T) {
		// DBに接続
		conn := NewConnectToDB()
		db, err := conn.Connect()
		defer db.Close()
		if err != nil {
			errors.New("Connect to DB failed.")
		}

		// DBの登録数を取得
		var count = 0
		// db.Model(&users.User{}).Count(&count)
		db.Table("users").Count(&count)

		if count != 0 {
			t.Errorf("record count invalid: got %v want %v",
				count, 0)
		}
	})
}
