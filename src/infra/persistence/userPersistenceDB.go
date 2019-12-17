package persistence

import (
	"context"

	_ "github.com/go-sql-driver/mysql"

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
	conn := NewConnectToDB()
	db, err := conn.Connect()
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	user1 := users.User{
		Name:  "namae",
		Email: "user1@example.com",
	}

	return []*users.User{&user1}, nil
}

// CreateUser ユーザー作成
func (userDB UserPersistanceDB) CreateUser(users.User) error {
	return nil
}
