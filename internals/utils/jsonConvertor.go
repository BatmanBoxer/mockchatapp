package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, object any) error {
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(object)
}

func ReadJson(r *http.Request, object interface{}) error {
	err := json.NewDecoder(r.Body).Decode(object)
	if err != nil {
		return err
	}
	return nil
}

