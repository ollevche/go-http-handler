package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	h := ExampleHandler{}

	h.RegisterRoutes(r)

	http.ListenAndServe(":8080", r)
}
