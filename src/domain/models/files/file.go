package files

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// File info
type File struct {
	ID   int64  `gorm:"column:id;primary_key"`
	Name string `gorm:"column:name" sql:"not null;type:varchar(256)"`
	Data []byte `gorm:"column:data" sql:"not null;type:mediumblob`
}
