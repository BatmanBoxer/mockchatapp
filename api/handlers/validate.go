package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/batmanboxer/mockchatapp/internals/authentication"
)

func(hanlders Handlers) ValidateHanlder(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
		return nil
	}
	body := struct {
		Jwt string `json:"jwt"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("unknown data type provided"))
		return nil
	}

	data, err := auth.ValidateJwt(body.Jwt)

	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("invalid Jwt"))
		return nil
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(data))
	return nil
}
