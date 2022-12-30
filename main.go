package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/psanti93/lenslockedv2.com/controllers"
	"github.com/psanti93/lenslockedv2.com/models"
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
	///////////////////////////////////////////////////////////////////////////////////////////////////

	// User Example

	// Set up connection to db
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	err = db.Ping()
	fmt.Println("Pinging from main application")
	if err != nil {
		panic(err)
	}

	//Set user service
	userService := models.UserService{
		DB: db,
	}

	// users controller
	usersC := controllers.Users{
		UserService: &userService,
	}
	usersC.Templates.New = views.Must(views.ParseFs(
		templates.FS,
		"signup.gohtml",
		"tailwind.gohtml",
	))

	usersC.Templates.SignIn = views.Must(views.ParseFs(
		templates.FS,
		"sigin.gohtml",
		"tailwind.gohtml",
	))

	r.Get("/signup", usersC.NewUser)
	r.Post("/signup", usersC.CreateUser)
	r.Get("/signin", usersC.SignIn)

	//////////////////////////////////////////////////////////////////////////////////////////

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)

	})
	fmt.Println("Starting the server on :3000...")

	http.ListenAndServe(":3000", r)
}
