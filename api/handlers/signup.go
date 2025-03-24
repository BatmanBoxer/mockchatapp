package handlers

import (
	"errors"
	"net/http"
	"regexp"
	"strings"
	"github.com/batmanboxer/mockchatapp/common"
	"github.com/batmanboxer/mockchatapp/internals/utils"
	"github.com/batmanboxer/mockchatapp/models"
)

func (handler *Handlers) SignUpHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
		return nil
	}

	data := models.SignUpData{}
	err := utils.ReadJson(r, &data)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Wrong Format Request Body"))
		return nil
	}

	err = ValidateName(&data.Name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Name is Invalid"))
		return nil
	}

	err = ValidateEmail(&data.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Email is Invalid"))
		return nil
	}
     
	err = handler.auth.AuthSignUp(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to Sign Up"))
		return err
	}

	status := models.SignUpSucess{Status: common.SUCCESS}
	utils.WriteJson(w, status)
	return nil
}

func ValidateName(name *string) error {
	if name == nil || strings.TrimSpace(*name) == "" {
		return errors.New("name cannot be empty")
	}
	if len(*name) < 3 {
		return errors.New("name must be at least 3 characters long")
	}
	return nil
}

func ValidateEmail(email *string) error {
	if email == nil || strings.TrimSpace(*email) == "" {
		return errors.New("email cannot be empty")
	}
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	matched, err := regexp.MatchString(emailRegex, *email)
	if err != nil {
		return err
	}

	if !matched {
		return errors.New("invalid email format")
	}

	return nil
}
