package handlers

import (
	"context"
	"github.com/batmanboxer/mockchatapp/common"
	"github.com/batmanboxer/mockchatapp/internals/authentication"
	"github.com/batmanboxer/mockchatapp/models"
	"log"
	"net/http"
	"sync"
)

type WebsocketStorage interface {
	GetMessages(string, int, int) ([]models.MessageModel, error)
	AddMessage(messageModel models.MessageModel) error
}

type Handlers struct {
	websocketStorage WebsocketStorage
	auth             *auth.Auth
	client           map[string][]*models.Client
	mutex            *sync.RWMutex
}

func NewHandlers(
	websocketStorage WebsocketStorage,
	auth *auth.Auth,
	client map[string][]*models.Client,
	mutex *sync.RWMutex,
) *Handlers {
	return &Handlers{
		websocketStorage: websocketStorage,
		auth:             auth,
		client:           client,
		mutex:            mutex,
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
    //also check if this user exists in userdatabase

		ctx := context.WithValue(r.Context(), common.CONTEXTIDKEY, userId)

		next(w, r.WithContext(ctx))
	}
}

