package handlers

import (
	"log"
	"net/http"
	"sync"

	"github.com/batmanboxer/mockchatapp/internals/authentication"
	"github.com/batmanboxer/mockchatapp/internals/database"
	"github.com/batmanboxer/mockchatapp/models"
)

type Handlers struct {
	db   database.Storage
	auth *auth.Auth
	client map[string][]*models.Client
  mutex *sync.RWMutex
}

func NewHandlers(
	db database.Storage,
	auth *auth.Auth,
	client map[string][]*models.Client,
  mutex *sync.RWMutex,
) *Handlers {
	return &Handlers{
		db:   db,
		auth: auth,
		client: client,
    mutex: mutex,
	}
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
