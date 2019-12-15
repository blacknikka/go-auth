package usecases

import (
	"context"

	"github.com/blacknikka/go-auth/domain/models/users"
	"github.com/blacknikka/go-auth/domain/repositories"
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
