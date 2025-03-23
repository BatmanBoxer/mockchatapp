package handlers

import (
	"log"
	"net/http"
	"github.com/batmanboxer/mockchatapp/internals/database"
)

type Handlers struct{
  db database.Storage
}

func NewHandlers(db database.Storage)*Handlers{
  return &Handlers{db: db}
}

type customHttpHandler func(http.ResponseWriter, *http.Request) error

func (h *Handlers)WrapperHandler(customHandler customHttpHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := customHandler(w, r)
		if err != nil {
			log.Printf("error: %s\n", err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}


