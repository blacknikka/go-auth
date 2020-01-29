package repositories

import (
	"context"

	"github.com/blacknikka/go-auth/domain/models/users"
)

// UserRepository Userのリポジトリのインターフェース
type UserRepository interface {
	GetAll(context.Context) ([]users.User, error)
	CreateUser(users.User) error
}
