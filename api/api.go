package api

import (
	"net/http"
	"github.com/batmanboxer/mockchatapp/api/handlers"
	"github.com/batmanboxer/mockchatapp/internals/database"
	"github.com/gorilla/mux"
)

type Api struct{
  port string;
  storage database.Storage
}

func NewApi(port string,storage database.Storage)*Api{
  return &Api{
    port: port,
    storage: storage,
  }
}

func (api *Api)StartApi(){
  handlers := handlers.NewHandlers(api.storage)
	mux := mux.NewRouter()
	mux.HandleFunc("/login", handlers.WrapperHandler(handlers.LoginHandler))
	mux.HandleFunc("/signup", handlers.WrapperHandler(handlers.SignUpHandler))
	mux.HandleFunc("/validate", handlers.WrapperHandler(handlers.ValidateHanlder))

	http.ListenAndServe(":4000", mux)

}
