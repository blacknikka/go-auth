package persistence

import (
	"context"

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

	user1 := users.User{
		Name:  "namae",
		Email: users.Email{Email: "user1@example.com"},
	}

	return []*users.User{&user1}, nil
}