package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/batmanboxer/mockchatapp/internals/utils"
	"github.com/batmanboxer/mockchatapp/models"
)

func(handlers *Handlers) LoginHandler(w http.ResponseWriter, r *http.Request) error {
  
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
		return nil
	}
	data := models.LoginData{}
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unknown Data Type Provided"))
		return nil
	}

	jwt, err := handlers.authManager.AuthLogin(data)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return nil
	}

	w.WriteHeader(http.StatusOK)
	sucess := models.LoginSucess{
		Jwt: jwt,
	}
  err = utils.WriteJson(w,sucess)
  
  if err != nil{
    return nil 
  }
	return nil
}

