package main

import (
	"log"
	"net/http"

	"github.com/blacknikka/go-auth/handlers"
)

func main() {
	handlers.InitializeRouting()
	log.Fatal(http.ListenAndServe(":5000", nil))
}
