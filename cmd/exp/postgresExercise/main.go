package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/psanti93/lenslockedv2.com/models"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBname   string
	SSLMode  string
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBname, cfg.SSLMode)
}

func main() {

	cfg := PostgresConfig{
		Host:     "localhost",
		Port:     "3355",
		User:     "baloo",
		Password: "junglebook",
		DBname:   "lenslockedv2",
		SSLMode:  "disable",
	}

	//pgx is the database driver
	db, err := sql.Open("pgx", cfg.String())

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

	user, err := us.Create("paul2@paul.com", "paul123")
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
