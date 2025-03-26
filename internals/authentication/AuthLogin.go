package auth

import (
	"errors"
	"github.com/batmanboxer/mockchatapp/models"
)

func (auth Auth) AuthLogin(data models.LoginData) (string, error) {
	account, err := auth.AuthDb.GetUserByEmail(data.Email)
	if err != nil {
		return "", err
	}
	if account.Password != data.Password {
		return "", errors.New("wrong password")
	}

	jwt, err := GenerateJwt(account.ID.String())
	if err != nil {
		return "", err
	}
	return jwt, nil
}
