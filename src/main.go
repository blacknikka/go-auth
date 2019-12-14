package main

import (
	"log"
	"net/http"

	"github.com/blacknikka/go-auth/controllers"
)

func main() {
	http.HandleFunc("/hello", controllers.HelloServer)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
