package main

import (
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello")
}

func main() {
	http.HandleFunc("/hello", HelloHandler)
	http.ListenAndServe(":12345", nil)
}
