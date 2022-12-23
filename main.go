package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/psanti93/lenslockedv2.com/controllers"
	"github.com/psanti93/lenslockedv2.com/templates"
	"github.com/psanti93/lenslockedv2.com/views"
)

func main() {
	r := chi.NewRouter()

	//parse the template before setting up handlers

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFs(
		templates.FS,
		"home.gohtml",
		"tailwind.gohtml",
	))))

	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFs(
		templates.FS,
		"contact.gohtml",
		"tailwind.gohtml",
	))))

	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFs(
		templates.FS,
		"faq.gohtml",
		"tailwind.gohtml",
	))))

	r.Get("/signUp", controllers.StaticHandler(views.Must(views.ParseFs(
		templates.FS,
		"signup.gohtml",
		"tailwind.gohtml",
	))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)

	})
	fmt.Println("Starting the server on :3000...")

	http.ListenAndServe(":3000", r)
}
