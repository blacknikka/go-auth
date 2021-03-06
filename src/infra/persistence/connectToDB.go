package persistence

import (
	"errors"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	// ErrDBOpen is an error message for opening the DB conneciton.
    ErrDBOpen = errors.New("When the DB connection is opened, an error is occurred")
)

// ConnectToDB DB接続
type ConnectToDB interface {
	Connect() (*gorm.DB, error)
}

// NewConnectToDB DB接続のstruct
func NewConnectToDB(dbConnection DBConnectionFactory) ConnectToDB {
	return &connectToDB{
		dbConnection: dbConnection,
	}
}

type connectToDB struct{
	dbConnection DBConnectionFactory
}

// Connect すべてを取得する
func (conn connectToDB) Connect() (*gorm.DB, error) {
	user := os.Getenv("DB_USER_NAME")
	password := os.Getenv("DB_USER_PASSWORD")
	host := "tcp(" + os.Getenv("DB_HOST") + ":3306)"
	dbTable := os.Getenv("MYSQL_DATABASE")
	connString := user + ":" + password + "@" + host + "/" + dbTable + "?charset=utf8&parseTime=True&loc=Local"

	db, err := conn.dbConnection.Open("mysql", connString)

	if err != nil {
		return nil, ErrDBOpen
	}

	return db, nil
}
