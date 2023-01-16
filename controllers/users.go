package controllers

import (
	"fmt"
	"net/http"

	"github.com/psanti93/lenslockedv2.com/models"
)

type Users struct {
	Templates struct {
		New    Template
		SignIn Template
	}
	UserService *models.UserService // connect the user service
}

func (u Users) NewUser(w http.ResponseWriter, r *http.Request) {

	var data struct {
		Email string
	}

	data.Email = r.FormValue("email")

	u.Templates.New.Execute(w, data)

}

func (u Users) CreateUser(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Persists imputs into the DB
	user, err := u.UserService.Create(email, password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User created: %+v ", user)

}

func (u Users) SignIn(w http.ResponseWriter, r *http.Request) {

	var data struct {
		Email string
	}

	data.Email = r.FormValue("email")

	u.Templates.SignIn.Execute(w, data)

}

func (u Users) ProcessSignIn(w http.ResponseWriter, r *http.Request) {

	var data struct {
		Email    string
		Password string
	}

	data.Email = r.FormValue("email")
	data.Password = r.FormValue("password")

	//Uses the DB to check and authenticate email and password
	user, err := u.UserService.Authenticate(data.Email, data.Password)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid log in", http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{
		Name:  "email",
		Value: user.Email,
		Path:  "/",
	}

	http.SetCookie(w, &cookie)

	fmt.Fprintf(w, "User authenticated: %+v", user)
}

func (u Users) CurrentUser(w http.ResponseWriter, r *http.Request) {
	email, err := r.Cookie("email")

	if err != nil {
		fmt.Errorf("cookie error: %w", err)
		return
	}

	fmt.Fprintf(w, "Email cookie:%s\n", email.Value)
	fmt.Fprintf(w, "Headers: %+v\n", r.Header["Cookie"]) //view the header value for cookie. Header is a giant map of headers from the request

}
