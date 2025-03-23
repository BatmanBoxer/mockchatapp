package database

import "github.com/batmanboxer/mockchatapp/models"

type Storage interface {
   AddAccount(models.SignUpData)error;
}
