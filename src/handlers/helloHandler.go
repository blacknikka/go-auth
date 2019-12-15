package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// HelloServer サンプル
func HelloServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world!")

	fmt.Println("hello world.")
}

// JSONRequest struct for request from client
type JSONRequest struct {
	Name string
}

// JSONResponse struct for response to client
type JSONResponse struct {
	Message string
}

// HelloJSONHandle json sample
func HelloJSONHandle(w http.ResponseWriter, r *http.Request) {
	body := JSONRequest{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	b, err := json.Marshal(JSONResponse{Message: "hello " + body.Name})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(b))
}
