package main

import (
	"github.com/batmanboxer/mockchatapp/api"
	"github.com/batmanboxer/mockchatapp/internals/database/postgress"
	"log"
)

func main() {
	postges, err := postgress.NewPostGres()

	if err != nil {
		log.Fatal("unable to make connection to database")
	}

	api := api.NewApi(":4000", postges)
	api.StartApi()
}
