package main

import (
	"log"
	"net/http"

	"github.com/blacknikka/go-auth/controllers"
)

func main() {
	controllers.InitializeRouting()
	log.Fatal(http.ListenAndServe(":5000", nil))
}
