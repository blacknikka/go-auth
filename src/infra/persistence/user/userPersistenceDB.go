package user

import (
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/blacknikka/go-auth/domain/models/users"
	"github.com/blacknikka/go-auth/domain/repositories"
)

// UserPersistanceDB DBでの永続化を行う
type UserPersistanceDB struct {
	db *gorm.DB
}

// NewUserPersistence 永続化Objectを返す
func NewUserPersistence(db *gorm.DB) repositories.UserRepository {
	return &UserPersistanceDB{
		db: db,
	}
}

// GetAll すべてを取得する
func (userDB UserPersistanceDB) GetAll() (*[]users.User, error) {
	users := []users.User{}
	userDB.db.Find(&users)

	return &users, nil
}

// FindByID IdからUserを取得する
func (userDB UserPersistanceDB) FindByID(id int) (*users.User, error) {
	user := users.User{}
	userDB.db.First(&user, id)

	if user.ID <= 0 {
		return nil, errors.New("User not found")
	}

	return &user, nil
}

// CreateUser ユーザー作成
func (userDB UserPersistanceDB) CreateUser(user users.User) (*users.User, error) {
	if result := userDB.db.NewRecord(user); result == false {
		return nil, errors.New("create user failed")
	}
	userDB.db.Create(&user)
	return &user, nil
}

// UpdateUser ユーザー更新
func (userDB UserPersistanceDB) UpdateUser(user users.User) (*users.User, error) {
	if user.ID <= 0 {
		return nil, errors.New("user ID is invalid")
	}

	userForFind := user
	if err := userDB.db.First(&userForFind).Error; err != nil {
		// not exists
		return nil, errors.New("ID doesn't exist")
	}

	userDB.db.Model(&user).Updates(
		users.User{
			Name:  user.Name,
			Email: user.Email,
		})

	return &user, nil
}

// DeleteUser ユーザー削除
func (userDB UserPersistanceDB) DeleteUser(user users.User) error {
	if user.ID <= 0 {
		return errors.New("user ID is invalid")
	}

	userForFind := user
	if err := userDB.db.First(&userForFind).Error; err != nil {
		// not exists
		return errors.New("ID doesn't exist")
	}

	if err := userDB.db.Model(&user).Delete(&user).Error; err != nil {
		return errors.New("Delete failed")
	}

	return nil
}
