package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	h := handler.NewSimpleHandler(apiClient)
	r := chi.NewRouter()
	r.Get("/hello", h.Hello)

	path := cfg.Host + ":" + cfg.Port

	srv := http.Server{
		Addr:    path,
		Handler: r,
	}

	// gracefull shutdown
	quit := make(chan os.Signal, 1)
	done := make(chan error, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		err := srv.Shutdown(ctx)
		done <- err
	}()

	log.Println("Starting server")

	_ = srv.ListenAndServe()

	err = <-done
	log.Println("Server shutting down with ", err)
}
