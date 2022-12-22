package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/psanti93/lenslockedv2.com/views"
)

func executeTemplate(w http.ResponseWriter, filepath string) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl, err := template.ParseFiles(filepath)

	if err != nil {
		log.Printf("Parsing template: %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}

	viewTmpl := views.Template{
		HTMLTmpl: tmpl,
	}

	viewTmpl.Execute(w, nil)

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tmplPath)

}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tmplPath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("templates", "faq.gohtml")
	executeTemplate(w, tmplPath)
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)

	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
