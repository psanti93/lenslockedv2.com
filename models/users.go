package models

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/psanti93/lenslockedv2.com/utils"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int
	Email        string
	PasswordHash string
}

type UserService struct {
	DB *sql.DB
}

func (us *UserService) Create(email string, password string) (*User, error) {
	// prep email and password to be stored in DB
	email = strings.ToLower(email) //consistently have emails be lower case when passing in
	passwordHash, err := utils.Hash(password)

	if err != nil {
		return nil, fmt.Errorf("Password for user couldn't be hashed: %w", err)
	}

	// prepping user object to be returned
	user := User{
		Email:        email,
		PasswordHash: passwordHash,
	}

	// Insert into DB
	row := us.DB.QueryRow(`
		INSERT INTO users(email, password_hash)
		VALUES($1,$2) RETURNING id;
	`, email, passwordHash)

	err = row.Scan(&user.ID) // get ID back from DB to use as value into user object on line 30

	if err != nil {
		return nil, fmt.Errorf("Create users: %w", err)
	}

	return &user, nil
}

func (userService *UserService) Authenticate(email string, password string) (*User, error) {
	email = strings.ToLower(email)
	user := User{
		Email: email,
	}

	// grabbing the user based on the email
	row := userService.DB.QueryRow(`
		SELECT id, password_hash 
		FROM users WHERE email=$1
	`, email)

	// returning it's id and password hash
	err := row.Scan(&user.ID, &user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("Authenticate: %w", err)
	}

	// where the authentication logic really happens: makes sure that the password hash and password provided match
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("Authenticate: %w", err)
	}

	return &user, nil
}
