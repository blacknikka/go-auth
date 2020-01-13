package usecases

import (
	"context"
	"errors"

	"github.com/blacknikka/go-auth/domain/models/users"
	"github.com/blacknikka/go-auth/domain/repositories"
	"github.com/blacknikka/go-auth/infra/persistence"
)

var (
	// ErrConnectingToDB is an error message for connection failed.
    ErrConnectingToDB = errors.New("Connect to DB failed.")
)

// UserUseCase Userのユースケース
type UserUseCase interface {
	GetAll(context.Context) ([]*users.User, error)
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
func (uu userUseCase) GetAll(ctx context.Context) ([]*users.User, error) {
	users, err := uu.userRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// CreateUser ユーザを作成する
func (uu userUseCase) CreateUser(user users.User) error {
	conn := persistence.NewConnectToDB()
	db, err := conn.Connect()
	defer db.Close()
	if err != nil {
		return ErrConnectingToDB
	}

	return uu.userRepository.CreateUser(db, user)
}
