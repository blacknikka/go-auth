package handlers

import (
	"net/http"

	"github.com/blacknikka/go-auth/handlers/rest"
	"github.com/blacknikka/go-auth/infra/persistence"
	"github.com/blacknikka/go-auth/usecases"
)

// InitializeRouting routing
func InitializeRouting() {
	userPersistence := persistence.NewUserPersistence()
	userUseCase := usecases.UserUseCase(userPersistence)
	userHandler := rest.NewUserHandler(userUseCase)

	http.HandleFunc("/hello", HelloServer)
	http.HandleFunc("/user", userHandler.Index)
}
