package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Contact Page</h1><p> to get in touch email me at <a href=\"mailto:paulsantiago282@gmail.com\">paulsantiago282@gmail.com</a>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<h1>FAQ Page</h1>
	<ul>
		<li>
			<b> Is there a free version?</b>
			No there is Not!
		</li>
		<li>
			<b> How many courses are there?</b>
			  There are 23 differnt courses
		</li>
		<li>
			<b> How many licks does it take to get to tootsie pop?</b>
			364
		</li>
	</ul>
	`)
}

// exercise 1:
func retreiveUserHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")

	w.Write([]byte(fmt.Sprintf("Hello %v", userId)))
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger) //excercise 2: adding the logger middleware
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.Get("/user/{userId}", retreiveUserHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)

	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
