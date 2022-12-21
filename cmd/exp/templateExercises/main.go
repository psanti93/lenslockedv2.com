package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/go-chi/chi"
)

// exercises 1 and 2
type User struct {
	Name     string
	Age      int
	IsPerson bool
	Children []string
	Meta     MetaData
}

type Children struct {
	Name string
}

type MetaData struct {
	Vists  int
	Cities map[string]string
	Pet    string
}

func expHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl, err := template.ParseFiles("templateEx.gohtml")

	if err != nil {
		log.Printf("Error Parsing template: %v", err)
		http.Error(w, "Template could not be parsed", http.StatusInternalServerError)
		return

	}

	//slice
	children := []string{"Michael", "Cora", "Alaina"}

	//map
	cities := map[string]string{"France": "Paris", "England": "London"}

	user := User{
		Name:     "Paul",
		Age:      23,
		IsPerson: true,
		Children: children,
		Meta: MetaData{
			Vists:  3,
			Cities: cities,
			Pet:    "Chewy",
		},
	}

	if err = tmpl.Execute(w, user); err != nil {
		log.Printf("Error executing template: %v ", err)
		http.Error(w, "Template could not be Executed", http.StatusInternalServerError)
		return
	}

}

func main() {

	r := chi.NewRouter()
	r.HandleFunc("/exp", expHandler)

	http.ListenAndServe(":3000", r)

}
