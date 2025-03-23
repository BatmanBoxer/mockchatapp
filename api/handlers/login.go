package handlers

import (
	"encoding/json"
	auth "github.com/batmanboxer/mockchatapp/internals/authentication"
	"github.com/batmanboxer/mockchatapp/models"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) error {
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

	jwt, err := auth.AuthLogin(data)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return nil
	}

	w.WriteHeader(http.StatusOK)
	sucess := models.LoginSucess{
		Jwt: jwt,
	}

	json.NewEncoder(w).Encode(sucess)
	return nil
}
