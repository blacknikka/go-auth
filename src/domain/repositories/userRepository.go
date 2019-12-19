package repositories

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/blacknikka/go-auth/domain/models/users"
)

// UserRepository Userのリポジトリのインターフェース
type UserRepository interface {
	GetAll(context.Context) ([]*users.User, error)
	CreateUser(*gorm.DB, users.User) error
}
