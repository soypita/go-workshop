package handler

import (
	"fmt"
	"net/http"

	"github.com/soypita/go-workshop/internals/api"
)

type SimpleHandler struct {
	jokeClient api.Client
}

func NewSimpleHandler(client api.Client) *SimpleHandler {
	return &SimpleHandler{
		jokeClient: client,
	}
}

func (h *SimpleHandler) Hello(w http.ResponseWriter, r *http.Request) {
	joke, err := h.jokeClient.GetJoke()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, joke.Joke)
}
