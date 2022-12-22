package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/psanti93/lenslockedv2.com/controllers"
	"github.com/psanti93/lenslockedv2.com/views"
)

func main() {
	r := chi.NewRouter()

	//parse the template before setting up handlers

	tmpl, err := views.Parse(filepath.Join("templates", "home.gohtml"))

	if err != nil {
		panic(err)
	}

	r.Get("/", controllers.StaticHandler(tmpl))

	tmpl, err = views.Parse(filepath.Join("templates", "contact.gohtml"))

	if err != nil {
		panic(err)
	}

	r.Get("/contact", controllers.StaticHandler(tmpl))

	tmpl, err = views.Parse(filepath.Join("templates", "faq.gohtml"))

	r.Get("/faq", controllers.StaticHandler(tmpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)

	})
	fmt.Println("Starting the server on :3000...")

	http.ListenAndServe(":3000", r)
}
