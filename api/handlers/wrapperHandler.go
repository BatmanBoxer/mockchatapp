package handlers

import (
	"log"
	"net/http"
	"sync"

	"github.com/batmanboxer/mockchatapp/internals/authentication"
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn    *websocket.Conn
	Message chan string 
}
type Handlers struct {
	auth *auth.Auth
  conn *map[string]*Client
  mutex *sync.RWMutex
}

func NewHandlers(auth *auth.Auth,conn *map[string]*Client, mutex *sync.RWMutex) *Handlers {
	return &Handlers{
    auth:auth,
    conn: conn,
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
