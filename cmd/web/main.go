package main

import (
	"github.com/nitishfy/go-webdev/pkg/handler"
	"net/http"
)

const portNum = ":8080"

func main() {
	http.HandleFunc("/home", handler.Home)
	http.HandleFunc("/about", handler.About)
	http.ListenAndServe(portNum, nil)
}
