package persistence

import (
	"testing"
)

func TestNewConnectToDB(t *testing.T) {
	t.Run("NewConnectToDB", func(t *testing.T) {
		connectToDB := NewConnectToDB()

		if connectToDB == nil {
			t.Error("NewConnectToDB() should return NOT nil value.")
		}
	})
}
