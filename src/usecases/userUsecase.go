package usecases

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
	GetAll() ([]users.User, error)
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
func (uu userUseCase) GetAll() ([]users.User, error) {
	users, err := uu.userRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// CreateUser ユーザを作成する
func (uu userUseCase) CreateUser(user users.User) (uses.User, error) {
	return uu.userRepository.CreateUser(user)
}
