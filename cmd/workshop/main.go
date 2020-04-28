package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ilyakaznacheev/cleanenv"

	"github.com/soypita/go-workshop/internals/api/jokes"
	"github.com/soypita/go-workshop/internals/config"
	"github.com/soypita/go-workshop/internals/handler"
)

func main() {
	cfg := config.ServerConfig{}
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	apiClient := jokes.NewJokeClient(cfg.JokeURL)

	h := handler.NewSimpleHandler()
	r := chi.NewRouter()
	r.Get("/hello", h.Hello)

	path := cfg.Host + ":" + cfg.Port
	log.Println("Starting server")

	err = http.ListenAndServe(path, r)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server shutting down")
}
