package api

import (
	"net/http"
	"sync"

	"github.com/batmanboxer/mockchatapp/api/handlers"
	"github.com/batmanboxer/mockchatapp/internals/authentication"
	"github.com/batmanboxer/mockchatapp/internals/database"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Api struct {
	port    string
	storage database.Storage
	conn    map[string]*websocket.Conn
	mutex   *sync.RWMutex
}

func NewApi(port string, storage database.Storage) *Api {
	return &Api{
		port:    port,
		storage: storage,
		conn:    make(map[string]*websocket.Conn),
		mutex:   &sync.RWMutex{},
	}
}

func (api *Api) StartApi() {
	handlers := handlers.NewHandlers(
		&auth.Auth{Db: api.storage},
		&api.conn,
		api.mutex,
	)
	mux := mux.NewRouter()

	mux.HandleFunc("/login", handlers.WrapperHandler(handlers.LoginHandler))
	mux.HandleFunc("/signup", handlers.WrapperHandler(handlers.SignUpHandler))
	mux.HandleFunc("/validate", handlers.WrapperHandler(handlers.ValidateHanlder))

	http.ListenAndServe(":4000", mux)
}
