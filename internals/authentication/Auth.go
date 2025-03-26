package auth

import (
	"github.com/batmanboxer/mockchatapp/models"
)

type AuthStorage interface {
	AddAccount(models.SignUpData) error
	GetUserByEmail(string) (models.AccountModel, error)
}
type Auth struct {
	AuthDb AuthStorage
}
