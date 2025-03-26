package api

import (
	"github.com/batmanboxer/mockchatapp/api/handlers"
	"github.com/batmanboxer/mockchatapp/internals/authentication"
	"github.com/batmanboxer/mockchatapp/models"
	"github.com/gorilla/mux"
	"net/http"
	"sync"
)

type Storage interface {
	AddAccount(models.SignUpData) error
	GetUserByEmail(string) (models.AccountModel, error)
	GetMessages(string, int, int) ([]models.MessageModel, error)
  AddMessage(messageModel models.MessageModel) error
}

type Api struct {
	port    string
	storage Storage
	conn    map[string][]*models.Client
	mutex   *sync.RWMutex
}

func NewApi(port string, storage Storage) *Api {
	return &Api{
		port:    port,
		storage: storage,
		conn:    make(map[string][]*models.Client),
		mutex:   &sync.RWMutex{},
	}
}

func (api *Api) StartApi() {
	auth := auth.Auth{
		AuthDb: api.storage,
	}

	handlers := handlers.NewHandlers(
		api.storage,
		&auth,
		api.conn,
		api.mutex,
	)

	mux := mux.NewRouter()

	mux.HandleFunc("/login", handlers.WrapperHandler(handlers.LoginHandler))
	mux.HandleFunc("/signup", handlers.WrapperHandler(handlers.SignUpHandler))
	mux.HandleFunc("/validate", handlers.WrapperHandler(handlers.ValidateHanlder))
	mux.HandleFunc("/listen/{id}", handlers.AuthenticationMiddleware(handlers.WrapperHandler(handlers.WebsocketHandler)))

	http.ListenAndServe(":4000", mux)
}
