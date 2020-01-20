package persistence

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DBConnectionFactory is a factory for DB connection
type DBConnectionFactory interface {
	Open(hostName string, connString string) (*gorm.DB, error)
}

type dBConnectionFactory struct{}

func NewDBConnectionFactory() DBConnectionFactory {
	return &dBConnectionFactory{}
}

func (factory *dBConnectionFactory) Open(hostName string, connString string) (*gorm.DB, error) {
	db, err := gorm.Open(
		hostName,
		connString)

	return db, err
}
