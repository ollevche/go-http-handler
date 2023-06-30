package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"reflect"
)

const reqBodyKey = "requestBody"

func JSONBody(destination any) func(http.Handler) http.Handler {
	if destination == nil {
		log.Fatal("BUG: non-nil destination value expected")
	}

	destinationType := reflect.TypeOf(destination)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Body == nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			destination := reflect.New(destinationType).Interface()

			err := json.NewDecoder(r.Body).Decode(&destination)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			ctx := context.WithValue(r.Context(), reqBodyKey, destination)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetReqBodyJSON(r *http.Request) any {
	return r.Context().Value(reqBodyKey)
}
