package controllers

import (
	"net/http"

	"github.com/psanti93/lenslockedv2.com/views"
)

type Users struct {
	Templates struct {
		New views.Template
	}
}

func (u Users) NewUser(w http.ResponseWriter, r *http.Request) {

	u.Templates.New.Execute(w, nil)

}
