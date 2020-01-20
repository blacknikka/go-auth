package persistence

import (
	"errors"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
)

// for mock
type fakeDBConnectionFactory struct {
	FakeOpen func(hostName string, connString string) (*gorm.DB, error)
}

func (s *fakeDBConnectionFactory) Open(hostName string, connString string) (*gorm.DB, error) {
	return s.FakeOpen(hostName, connString)
}

func TestConnectToDB(t *testing.T) {
	t.Run("NewConnectToDB", func(t *testing.T) {
		// DI (DBConnectionFactory)
		spy := &fakeDBConnectionFactory{
			FakeOpen: func(hostName string, connString string) (*gorm.DB, error) {
				// 適当なエラー
				return &gorm.DB{}, nil
			},
		}

		connectToDB := NewConnectToDB(spy)

		if connectToDB == nil {
			t.Error("NewConnectToDB() should return NOT nil value.")
		}
	})

	t.Run("Connect failed", func(t *testing.T) {
		// DI (UserUsDBConnectionFactoryeCase)
		spy := &fakeDBConnectionFactory{
			FakeOpen: func(hostName string, connString string) (*gorm.DB, error) {
				// 適当なエラー
				return nil, errors.New("connection error")
			},
		}

		connectToDB := NewConnectToDB(spy)
		db, err := connectToDB.Connect()

		if db != nil {
			t.Error("In this case, connectToDB.Connect() should return nil value of *gorm.DB")
		}

		if err != ErrDBOpen {
			t.Errorf("got error %q want %q", err, ErrDBOpen)
		}
	})
}
