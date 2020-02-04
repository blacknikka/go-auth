package repositories

import (
	"github.com/blacknikka/go-auth/domain/models/users"
)

// UserRepository Userのリポジトリのインターフェース
type UserRepository interface {
	GetAll() (*[]users.User, error)
	CreateUser(users.User) (*users.User, error)
	FindByID(int) (*users.User, error)
	UpdateUser(users.User) (*users.User, error)
	DeleteUser(users.User) (error)
}
