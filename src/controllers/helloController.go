package controllers

import (
	"fmt"
	"io"
	"net/http"
)

// HelloServer サンプル
func HelloServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world!")

	fmt.Println("hello world.")
}
