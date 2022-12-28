package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
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

	//CREATE a table..
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
		  id SERIAL PRIMARY KEY,
		  name TEXT,
		  email TEXT UNIQUE NOT NULL
		);

		CREATE TABLE IF NOT EXISTS orders (
		  id SERIAL PRIMARY KEY,
		  user_id INT NOT NULL,
		  amount INT,
		  description TEXT

		);
	`)

	if err != nil {
		panic(err)
	}

	fmt.Println("Tables Created")

	//Creating Sample Orders
	userId := 1

	for i := 1; i <= 5; i++ {
		amount := i * 100
		desc := fmt.Sprintf("Fake order #%d", i)
		_, err := db.Exec(`
			INSERT INTO orders(user_id,amount,description)
			VALUES($1,$2,$3)
		`, userId, amount, desc)

		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Fake orders created")

}
