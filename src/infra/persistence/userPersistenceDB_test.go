package persistence

import (
	"testing"

	"github.com/blacknikka/go-auth/domain/models/users"
	_ "github.com/go-sql-driver/mysql"
)

func TestDB(t *testing.T) {
	t.Run("CreateUser正常系", func(t *testing.T) {
		// DBに接続
		conn := NewConnectToDB()
		db, err := conn.Connect()
		defer db.Close()
		if err != nil {
			t.Fatal("Connect to DB failed.")
		}

		// DBの登録数を取得
		var count = 0
		db.Model(&users.User{}).Count(&count)
		// db.Table("users").Count(&count)

		if count != 0 {
			t.Errorf("record count invalid: got %v want %v",
				count, 0)
		}

		// CreateUser
		userDB := &UserPersistanceDB{}
		user := users.User{
			Name:  "user1",
			Email: "user1@example.com",
		}
		err = userDB.CreateUser(db, user)
		if err != nil {
			t.Error("insert error")
		}

		db.Model(&users.User{}).Count(&count)

		if count != 1 {
			t.Errorf("insert error: got %v want %v", count, 1)
		}
	})
}
