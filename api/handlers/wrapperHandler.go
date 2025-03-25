package handlers

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/batmanboxer/mockchatapp/common"
	"github.com/batmanboxer/mockchatapp/internals/authentication"
	"github.com/batmanboxer/mockchatapp/internals/database"
	"github.com/batmanboxer/mockchatapp/models"
)




type Handlers struct {
	db     database.Storage
	auth   *auth.Auth
	client map[string][]*models.Client
	mutex  *sync.RWMutex
}

func NewHandlers(
	db database.Storage,
	auth *auth.Auth,
	client map[string][]*models.Client,
	mutex *sync.RWMutex,
) *Handlers {
	return &Handlers{
		db:     db,
		auth:   auth,
		client: client,
		mutex:  mutex,
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

func (h *Handlers) AuthenticationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		userId, err := auth.ValidateJwt(authHeader)
		if err != nil {
			http.Error(w, "Invalid JWT", http.StatusUnauthorized)
		}
		ctx := context.WithValue(r.Context(), common.CONTEXTIDKEY, userId)

		next(w, r.WithContext(ctx))
	}
}
