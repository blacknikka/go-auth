package handlers

import (
	"net/http"

	"github.com/blacknikka/go-auth/handlers/rest"
	"github.com/blacknikka/go-auth/infra/persistence"
	"github.com/blacknikka/go-auth/usecases"
)

// InitializeRouting routing
func InitializeRouting() {
	conn := persistence.NewConnectToDB(persistence.NewDBConnectionFactory())
	db, err := conn.Connect()
	if err != nil {
		panic(err.Error())
	}

	userPersistence := persistence.NewUserPersistence(db)
	userUseCase := usecases.UserUseCase(userPersistence)
	userHandler := rest.NewUserHandler(userUseCase)

	http.HandleFunc("/hello", HelloServer)
	http.HandleFunc("/user", userHandler.Index)
}
