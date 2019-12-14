package controllers

import (
	"fmt"
	"io"
	"net/http"
)

// HelloServer サンプル
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")

	fmt.Println("hello world.")
}
