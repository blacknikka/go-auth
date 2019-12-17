package users

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User user info
type User struct {
	ID    int64  `gorm:"column:id;primary_key"`
	Name  string `gorm:"column:name" sql:"not null;type:varchar(256)"`
	Email string `gorm:"column:email" sql:"not null;type:varchar(256)"`
}
