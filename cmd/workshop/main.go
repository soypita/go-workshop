package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/soypita/go-workshop/internals/handler"
)

func main() {
	h := handler.NewSimpleHandler()
	r := chi.NewRouter()
	r.Get("/hello", h.Hello)

	log.Println("Starting server")
	err := http.ListenAndServe(":8080", r)
	log.Fatal(err)
	log.Println("Server shutting down")
}
