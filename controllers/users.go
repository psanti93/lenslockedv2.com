package controllers

import (
	"fmt"
	"net/http"
)

type Users struct {
	Templates struct {
		New Template
	}
}

func (u Users) NewUser(w http.ResponseWriter, r *http.Request) {

	u.Templates.New.Execute(w, nil)

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
