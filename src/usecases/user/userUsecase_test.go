package user

import (
	"errors"
	"testing"

	"github.com/blacknikka/go-auth/domain/models/users"
	"github.com/blacknikka/go-auth/domain/repositories"
)

// for mock
type fakeUserRepository struct {
	repositories.UserRepository
	FakeGetAll     func() (*[]users.User, error)
	FakeCreateUser func(users.User) (*users.User, error)
	FakeFindByID   func(int) (*users.User, error)
	FakeUpdateUser func(users.User) (*users.User, error)
	FakeDeleteUser func(users.User) error
}

func (s *fakeUserRepository) GetAll() (*[]users.User, error) {
	return s.FakeGetAll()
}

func (s *fakeUserRepository) CreateUser(user users.User) (*users.User, error) {
	return s.FakeCreateUser(user)
}

func (s *fakeUserRepository) FindByID(id int) (*users.User, error) {
	return s.FakeFindByID(id)
}

func (s *fakeUserRepository) UpdateUser(user users.User) (*users.User, error) {
	return s.FakeUpdateUser(user)
}

func (s *fakeUserRepository) DeleteUser(user users.User) error {
	return s.FakeDeleteUser(user)
}

func TestGetAll(t *testing.T) {
	t.Run("GetAll正常系", func(t *testing.T) {
		spyUserRepository := &fakeUserRepository{
			FakeGetAll: func() (*[]users.User, error) {
				return &[]users.User{
					users.User{},
					users.User{},
				}, nil
			},
		}

		usecase := NewUserUseCase(spyUserRepository)

		resultUsers, err := usecase.GetAll()

		if err != nil {
			t.Errorf("err should be nil: %v", err)
		}

		if resultUsers == nil {
			t.Errorf("Returned USERS shouldn't be nil: %v", resultUsers)
		}
	})

	t.Run("GetAll異常系", func(t *testing.T) {
		spyUserRepository := &fakeUserRepository{
			FakeGetAll: func() (*[]users.User, error) {
				return nil, errors.New("something error")
			},
		}

		usecase := NewUserUseCase(spyUserRepository)

		// 何かのエラーが発生する
		resultUsers, err := usecase.GetAll()

		if err == nil {
			t.Errorf("err shouldn't be nil: %v", err)
		}

		if resultUsers != nil {
			t.Errorf("Returned USERS should be nil: %v", resultUsers)
		}
	})
}
