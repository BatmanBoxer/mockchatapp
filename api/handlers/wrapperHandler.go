package handlers

import (
	"log"
	"net/http"
	"github.com/batmanboxer/mockchatapp/internals/authentication"
)

type Handlers struct {
	auth *auth.Auth
}

func NewHandlers(auth *auth.Auth) *Handlers {
	return &Handlers{auth:auth}
}

type customHttpHandler func(http.ResponseWriter, *http.Request) error

func (h *Handlers) WrapperHandler(customHandler customHttpHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := customHandler(w, r)
		if err != nil {
			log.Printf("error: %s\n", err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}
