package controllers

import (
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
