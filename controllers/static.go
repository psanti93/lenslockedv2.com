package controllers

import (
	"net/http"

	"github.com/psanti93/lenslockedv2.com/views"
)

func StaticHandler(tmpl views.Template) http.HandlerFunc {
	//closure example
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	}
}
