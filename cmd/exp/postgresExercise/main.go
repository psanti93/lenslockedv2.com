package main

import (
	"fmt"

	"github.com/psanti93/lenslockedv2.com/models"
)

func main() {

	// Testing Postgres Config in models
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)

	if err != nil {
		panic(err)
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err) // it will panic if the docker container is not spun up
	}
	fmt.Println("Connected!") //it's connected but you don't know if it's able to communicate with the DB

	// TESTING the user service (make users table first)
	us := models.UserService{
		DB: db,
	}

	user, err := us.Create("paul3@paul.com", "paul123")
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
