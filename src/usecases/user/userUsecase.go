package user

import (
	"errors"

	"github.com/blacknikka/go-auth/domain/models/users"
	"github.com/blacknikka/go-auth/domain/repositories"
)

var (
	// ErrConnectingToDB is an error message for connection failed.
	ErrConnectingToDB = errors.New("Connect to DB failed.")
)

// UserUseCase Userのユースケース
type UserUseCase interface {
	GetAll() (*[]users.User, error)
	CreateUser(users.User) (*users.User, error)
	FindByID(int) (*users.User, error)
	UpdateUser(users.User) (*users.User, error)
	DeleteUser(users.User) error
}

type userUseCase struct {
	userRepository repositories.UserRepository
}

// NewUserUseCase Userのユースケースを取得する
func NewUserUseCase(userRepository repositories.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}

// GetAll 全ユーザ取得する
func (uu userUseCase) GetAll() (*[]users.User, error) {
	users, err := uu.userRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// CreateUser ユーザを作成する
func (uu userUseCase) CreateUser(user users.User) (*users.User, error) {
	return uu.userRepository.CreateUser(user)
}

// CreateUser ユーザを探す
func (uu userUseCase) FindByID(id int) (*users.User, error) {
	return uu.userRepository.FindByID(id)
}

// UpdateUser ユーザ情報を更新
func (uu userUseCase) UpdateUser(user users.User) (*users.User, error) {
	return uu.userRepository.UpdateUser(user)
}

// DeleteUser ユーザ情報を削除
func (uu userUseCase) DeleteUser(user users.User) error {
	return uu.userRepository.DeleteUser(user)
}
