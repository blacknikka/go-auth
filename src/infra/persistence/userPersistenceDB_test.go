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
	teardown()

	os.Exit(ret)
}

func TestCreateUser(t *testing.T) {
	t.Run("CreateUser正常系", func(t *testing.T) {
		// DBを空にする
		db.Delete(&users.User{})

		// DBの登録数を取得
		var count = 0
		db.Model(&users.User{}).Count(&count)

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
		createdUser, err := userDB.CreateUser(user)
		if err != nil {
			t.Error("insert error")
		}

		db.Model(&users.User{}).Count(&count)

		if count != 1 {
			t.Errorf("insert error: got %v want %v", count, 1)
		}

		if createdUser.ID <= 0 {
			t.Errorf("created ID should be a positive value: %v", createdUser.ID)
		}
	})
}

func TestFindByID(t *testing.T) {
	t.Run("FindByID正常系", func(t *testing.T) {
		// DBを空にする
		db.Delete(&users.User{})

		// DBの登録数を取得
		var count = 0
		db.Model(&users.User{}).Count(&count)

		if count != 0 {
			t.Errorf("record count invalid: got %v want %v",
				count, 0)
		}

		// Userの登録
		userDB := &UserPersistanceDB{db: db}
		targetUser := users.User{
			Name:  "user1",
			Email: "user1@example.com",
		}
		createdUser, _ := userDB.CreateUser(targetUser)

		if createdUser.ID <= 0 {
			t.Errorf("created ID should be positive value: %v",
			createdUser.ID)
		}

		// FindByID
		// IDで検索を行う（上記登録したUserで検索）
		foundUser, err := userDB.FindByID(int(createdUser.ID))

		if createdUser == nil {
			t.Errorf("Found user shouldn't be nil: %v", foundUser)
		}

		if foundUser.ID != createdUser.ID {
			t.Errorf("found user is invalid: got %v want %v",
				foundUser, createdUser)
		}

		if foundUser.Name != targetUser.Name {
			t.Errorf("found user is invalid: got %v want %v", foundUser.Name, targetUser.Name)
		}

		if foundUser.Email != targetUser.Email {
			t.Errorf("found user is invalid: got %v want %v", foundUser.Email, targetUser.Email)
		}

		if err != nil {
			t.Errorf("err should be nil, but %v", err)
		}
	})

	t.Run("FindByID異常系", func(t *testing.T) {
		// DBを空にする
		db.Delete(&users.User{})

		// Userの登録
		userDB := &UserPersistanceDB{db: db}
		targetUser := users.User{
			Name:  "user1",
			Email: "user1@example.com",
		}
		createdUser, _ := userDB.CreateUser(targetUser)

		// find(存在しないIDを指定する)
		foundUser, err := userDB.FindByID(int(createdUser.ID + 1))

		if foundUser != nil {
			t.Errorf("Since the ID is invalid, user should be nil: %v", foundUser)
		}

		if err == nil {
			t.Errorf("Error shouldn't be nil: %v", err)
		}
	})

}
