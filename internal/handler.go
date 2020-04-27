package handler

import (
	"fmt"
	"net/http"
)

type SimpleHandler struct {
}

func NewSimpleHandler() *SimpleHandler {
	return &SimpleHandler{}
}

func (h *SimpleHandler) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello golang")
}
