package main

import (
	"net/http"
)


func main() {
	http.HandleFunc("/user", makeHTTPHandler(handleGetUserById))
	http.ListenAndServe(":8080", nil)
}
