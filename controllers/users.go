package controllers

import (
	"fmt"
	"net/http"

	"github.com/psanti93/lenslockedv2.com/models"
)

type Users struct {
	Templates struct {
		New Template
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
	if err := r.ParseForm(); err != nil {
		fmt.Errorf("Error in parsing request: %s:", err)
		http.Error(w, "There was an error parsing the request", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "<p> New User Added! </p>")
	fmt.Printf("The email address is: %s", r.FormValue("email"))
	fmt.Printf("The password is: %s", r.FormValue("password"))

}
