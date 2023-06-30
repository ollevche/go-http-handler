package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

type ExampleHandler struct{}

func (h ExampleHandler) RegisterRoutes(r *mux.Router) {

	// create a handler with requried middleware
	// it's easy to add multiple middlewares (5+) with the following syntax:
	postExampleJSON := alice.New().
		Append(JSONBody(jsonExampleReqBody{})).
		// Append(ValidateReqBody).
		// Append(AllowAdminsOnly).
		// Append(LogIncomingRequest).
		// etc.
		ThenFunc(h.PostExampleJSON)

	r.Handle("/examples/json", postExampleJSON).
		Methods(http.MethodPost)
}

type jsonExampleReqBody struct {
	Hello   string `json:"hello"`
	Name    string `json:"name"`
	Counter int    `json:"counter"`
}

func (h ExampleHandler) PostExampleJSON(w http.ResponseWriter, r *http.Request) {
	example := GetReqBodyJSON(r).(*jsonExampleReqBody) // problem: no compile time static type checks, solution: generics?

	if example.Name != "" {
		log.Default().Printf("PostExampleJSON called with %v name", example.Name)
	}

	log.Default().Println(example)
}
