package auth

import (
	"errors"
	"github.com/batmanboxer/mockchatapp/models"
)

func AuthLogin(data models.LoginData) (string, error) {
  //Login Login Resides Here
	if data.Email != "batman" {
		return "", errors.New("email not matched")
	}
	jwt, err := GenerateJwt(data.Email)
	if err != nil {
		return "", err
	}
	return jwt, nil
}
