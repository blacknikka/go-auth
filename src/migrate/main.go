package main

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/blacknikka/go-auth/domain/models/users"
	"github.com/blacknikka/go-auth/infra/persistence"
)

func main() {

	conn := persistence.NewConnectToDB()
	db, err := conn.Connect()
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Migrate
	db.AutoMigrate(&users.User{})
}
