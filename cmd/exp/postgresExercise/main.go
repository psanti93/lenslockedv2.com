package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {

	//pgx is the database driver
	db, err := sql.Open("pgx", "host=localhost port=5432 user=baloo password=junglebook dbname=lenslockedv2 sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err) // it will panic if the docker container is not spun up
	}
	fmt.Println("Connected!") //it's connected but you don't know if it's able to communicate with the DB
}
