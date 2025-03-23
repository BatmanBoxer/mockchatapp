package main

import (
	"github.com/batmanboxer/mockchatapp/api/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/login", handlers.WrapperHandler(handlers.LoginHandler))
	mux.HandleFunc("/validate", handlers.WrapperHandler(handlers.ValidateHanlder))

	http.ListenAndServe(":4000", mux)
}
