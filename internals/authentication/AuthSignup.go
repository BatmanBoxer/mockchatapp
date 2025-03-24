package auth

import "github.com/batmanboxer/mockchatapp/models"

func (auth Auth) AuthSignUp(signUpData models.SignUpData) error {
	error := auth.Db.AddAccount(signUpData)

	if error != nil {
		return error
	}
	return nil
}
