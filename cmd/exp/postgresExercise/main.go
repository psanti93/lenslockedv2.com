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
		Port:     "5432",
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
	name := "Paul Santiago"
	email := "paul@test.com"

	//returns a single item from the DB
	row := db.QueryRow(`
		INSERT INTO users(name,email)
		VALUES ($1, $2) RETURNING id;
	`, name, email)

	//row.Err() // checks for any errors that occur on the sql statement. however not needed since the error will checked in ro.Scan
	var id int
	// take the value we got from the DB and store it in the id variable in memory
	if err := row.Scan(&id); err != nil {
		panic(err)
	}

	fmt.Printf("User Created. Id = %v", id)

	/** RESULT

	lenslockedv2=# select * from users;
	id |            name             |     email
	----+-----------------------------+---------------
	1 | ',''); DROP TABLE users; -- | bob@test.com
	2 | Paul Santiago               | paul@test.com


		Connected!
		Tables Created
		User Created. Id = 2

	**/

}
