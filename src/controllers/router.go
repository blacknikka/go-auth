package controllers

import (
	"net/http"
)

// InitializeRouting routing
func InitializeRouting() {
	http.HandleFunc("/hello", HelloServer)
}
