package persistence

import (
	"context"
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/blacknikka/go-auth/domain/models/users"
	"github.com/blacknikka/go-auth/domain/repositories"
)

// UserPersistanceDB DBでの永続化を行う
type UserPersistanceDB struct{
	db *gorm.DB
}

// NewUserPersistence 永続化Objectを返す
func NewUserPersistence(db *gorm.DB) repositories.UserRepository {
	return &UserPersistanceDB{
		db: db,
	}
}

// GetAll すべてを取得する
func (userDB UserPersistanceDB) GetAll(context.Context) ([]users.User, error) {
	users := []users.User{}
	userDB.db.Find(&users)

	return users, nil
}

// CreateUser ユーザー作成
func (userDB UserPersistanceDB) CreateUser(user users.User) error {
	if result := userDB.db.NewRecord(user); result == false {
		return errors.New("create user failed")
	}
	userDB.db.Create(&user)
	return nil
}
