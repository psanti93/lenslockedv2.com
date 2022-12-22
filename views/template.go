package views

import (
	"html/template"
	"log"
	"net/http"
)

type Template struct {
	HTMLTmpl *template.Template
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if err := t.HTMLTmpl.Execute(w, data); err != nil {
		log.Printf("Executing template:%v", err)
		http.Error(w, "There was an error executing the template", http.StatusInternalServerError)
		return
	}

}
