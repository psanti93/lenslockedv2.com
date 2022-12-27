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
	// doing this will delete the table users
	name := "',''); DROP TABLE users; --"
	email := "bob@test.com"

	// query := fmt.Sprintf(`
	// 	INSERT INTO users (name, email)
	// 	VALUES ('%s','%s');
	// `, name, email)
	// fmt.Printf("Executing query: %s\n", query)

	// _, err = db.Exec(query)

	/**   RESULT

			Executing query:
	                INSERT INTO users (name, email)
	                VALUES ('','');; DROP TABLE users; --','bob@test.com');

			lenslockedv2=# select * from users;
			ERROR:  relation "users" does not exist
			LINE 1: select * from users;

		**/

	//whereas using variable placeholders ($1,$2) will prevent an injection like in the previous example

	_, err = db.Exec(`
		INSERT INTO users(name,email)
		VALUES ($1, $2);
	`, name, email)

	if err != nil {
		panic(err)
	}

	fmt.Println("User Created")

	/** RESULT
		lenslockedv2=# select * from users;
		id |            name             |    email
		----+-----------------------------+--------------
		1 | ',''); DROP TABLE users; -- | bob@test.com

	**/

}
