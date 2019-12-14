package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// HelloServer サンプル
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")

	fmt.Println("hello world.")
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
