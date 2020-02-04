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

func TestUpdateUser(t *testing.T) {
	t.Run("UpdateUser正常系", func(t *testing.T) {
		// DBを空にする
		db.Delete(&users.User{})

		// Userの登録
		userDB := &UserPersistanceDB{db: db}
		targetUser := users.User{
			Name:  "user1",
			Email: "user1@example.com",
		}
		createdUser, _ := userDB.CreateUser(targetUser)

		// User情報の更新
		targetUser.Name = "my-name"
		updatedUser, err := userDB.UpdateUser(*createdUser)

		if updatedUser.Name != "my-name" && 
			updatedUser.Email != "user1@example.com" {
			t.Errorf("Update failed: %v", updatedUser)
		}

		if err != nil {
			t.Errorf("err should be nil: %v", err)
		}

		updatedUser.Email = "my-name@example.com"
		updatedUser, err = userDB.UpdateUser(*updatedUser)

		if updatedUser.Name != "my-name" && 
			updatedUser.Email != "my-name@example.com" {
			t.Errorf("Update failed: %v", updatedUser)
		}

		if err != nil {
			t.Errorf("err should be nil: %v", err)
		}

	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("DeleteUser正常系", func(t *testing.T) {
		// DBを空にする
		db.Delete(&users.User{})

		// Userの登録
		userDB := &UserPersistanceDB{db: db}
		targetUser := users.User{
			Name:  "user1",
			Email: "user1@example.com",
		}
		createdUser, _ := userDB.CreateUser(targetUser)

		// DBの登録数を取得
		var prevCount = 0
		db.Model(&users.User{}).Count(&prevCount)
		
		// User情報の削除
		err := userDB.DeleteUser(*createdUser)
		
		if err != nil {
			t.Errorf("Delete failed: %v", err)
		}

		var afterCount = 0
		db.Model(&users.User{}).Count(&afterCount)

		confirmBeforeAndAfter(t, prevCount - 1, afterCount)
	})

	t.Run("DeleteUser異常系_IDがゼロ", func(t *testing.T) {
		// DBを空にする
		db.Delete(&users.User{})

		// Userの登録
		userDB := &UserPersistanceDB{db: db}
		targetUser := users.User{
			Name:  "user1",
			Email: "user1@example.com",
		}
		_, _ = userDB.CreateUser(targetUser)

		// DBの登録数を取得
		var prevCount = 0
		db.Model(&users.User{}).Count(&prevCount)

		// User情報の削除(IDが異常)
		err := userDB.DeleteUser(users.User{})
		
		if err == nil {
			t.Errorf("Deleting should fail if the ID is zero: %v", err)
		} else if err.Error() != "user ID is invalid" {
			t.Errorf("Error message is invalid: %v", err)
		}

		// DBの登録数が変わっていないことを確認
		var afterCount = 0
		db.Model(&users.User{}).Count(&afterCount)

		confirmBeforeAndAfter(t, prevCount, afterCount)
	})

	t.Run("DeleteUser異常系_IDが存在しない", func(t *testing.T) {
		// DBを空にする
		db.Delete(&users.User{})

		// Userの登録
		userDB := &UserPersistanceDB{db: db}
		targetUser := users.User{
			Name:  "user1",
			Email: "user1@example.com",
		}
		createduser, _ := userDB.CreateUser(targetUser)

		// DBの登録数を取得
		var prevCount = 0
		db.Model(&users.User{}).Count(&prevCount)

		// User情報の削除(IDが異常)
		createduser.ID++		// 存在しないIDを作成
		err := userDB.DeleteUser(*createduser)
		
		if err == nil {
			t.Errorf("Deleting should fail if the ID is invalid: %v", err)
		} else if err.Error() != "ID doesn't exist" {
			t.Errorf("Error message is invalid: %v", err)
		}

		// DBの登録数が変わっていないことを確認
		var afterCount = 0
		db.Model(&users.User{}).Count(&afterCount)

		confirmBeforeAndAfter(t, prevCount, afterCount)
	})
}

// confirmBeforeAndAfter compares previous value and after value.
func confirmBeforeAndAfter(t *testing.T, prev int, after int) {
	t.Helper()

	if prev != after {
		t.Errorf("DB record should be the same between before and after the Delete: before => %v, after => %v", prev, after)
	}
}
