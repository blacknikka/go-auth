package persistence

import (
	"context"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/blacknikka/go-auth/domain/models/users"
	"github.com/blacknikka/go-auth/domain/repositories"
)

// UserPersistanceDB DBでの永続化を行う
type UserPersistanceDB struct{}

// NewUserPersistence 永続化Objectを返す
func NewUserPersistence() repositories.UserRepository {
	return &UserPersistanceDB{}
}

// GetAll すべてを取得する
func (userDB UserPersistanceDB) GetAll(context.Context) ([]*users.User, error) {
	user := os.Getenv("DB_USER_NAME")
	password := os.Getenv("DB_USER_PASSWORD")
	host := "tcp(" + os.Getenv("DB_HOST") + ":3306)"
	dbTable := os.Getenv("MYSQL_DATABASE")
	connString := user + ":" + password + "@" + host + "/" + dbTable + "?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(
		"mysql",
		connString)
	defer db.Close()

	if err != nil {
		panic(err.Error())
	}

	user1 := users.User{
		Name:  "namae",
		Email: users.Email{Email: "user1@example.com"},
	}

	return []*users.User{&user1}, nil
}
