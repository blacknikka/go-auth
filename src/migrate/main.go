package main

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/blacknikka/go-auth/domain/models/files"
	"github.com/blacknikka/go-auth/domain/models/users"
	"github.com/blacknikka/go-auth/infra/persistence"
)

func main() {

	conn := persistence.NewConnectToDB(persistence.NewDBConnectionFactory())
	db, err := conn.Connect()
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Check if the table exists
	if db.HasTable(&users.User{}) {
		db.DropTable(&users.User{})
	}

	if db.HasTable(&files.File{}) {
		db.DropTable(&files.File{})
	}

	// Migrate
	db.AutoMigrate(&users.User{})
	db.AutoMigrate(&files.File{})
}
