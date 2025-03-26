package models

import (
	"github.com/google/uuid"
	"time"
)

type LoginSucess struct {
	Jwt string `json:"jwt"`
}

type SignUpSucess struct {
	Status string `json:"status"`
}

// model represents database
// data represents request and respose

type LoginData struct {
	Email    string
	Password string
}

type SignUpData struct {
	Name     string
	Age      int
	Email    string
	Password string
}
type AccountModel struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Verified  bool      `json:"verified"`
	CreatedAt time.Time `json:"created_at"`
}
