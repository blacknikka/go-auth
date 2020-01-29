package persistence

import (
	"os"
	"testing"

	"github.com/blacknikka/go-auth/domain/models/users"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func setup() {
	// DBに接続
	conn := NewConnectToDB(NewDBConnectionFactory())
	var err error
	db, err = conn.Connect()
	if err != nil {
		panic("Connect to DB failed.")
	}

	db.Delete(&users.User{})
}

func teardown() {
	defer db.Close()
}

func TestMain(m *testing.M) {
	setup()
	ret := m.Run()
	if ret == 0 {
		teardown()
	}
	os.Exit(ret)
}

func TestDB(t *testing.T) {
	t.Run("CreateUser正常系", func(t *testing.T) {
		// DBの登録数を取得
		var count = 0
		db.Model(&users.User{}).Count(&count)
		// db.Table("users").Count(&count)

		if count != 0 {
			t.Errorf("record count invalid: got %v want %v",
				count, 0)
		}

		// CreateUser
		userDB := &UserPersistanceDB{db: db}
		user := users.User{
			Name:  "user1",
			Email: "user1@example.com",
		}
		err := userDB.CreateUser(user)
		if err != nil {
			t.Error("insert error")
		}

		db.Model(&users.User{}).Count(&count)

		if count != 1 {
			t.Errorf("insert error: got %v want %v", count, 1)
		}
	})
}
