package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Contact Page</h1><p> to get in touch email me at <a href=\"mailto:paulsantiago282@gmail.com\">paulsantiago282@gmail.com</a>")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, "Page Not Found", http.StatusNotFound)
	}
}

func main() {
	var router http.HandlerFunc
	router = pathHandler
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", router) //ListenAndServe requires a handler type
	//http.ListenAndServe(":3000", http.HandlerFunc(pathHandler)) //casting pathHandler to type HandlerFunc
}

// in Go you can use functions like any other data type because of that you can methods to that type

// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers. If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler that calls f.

// type HandlerFunc func(ResponseWriter, *Request)

// // ServeHTTP calls f(w, r).
// func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
// 	f(w, r)
// }

// Difference
// http.Handler - interface with ServHTTP method
// http.HandlerFunc - function type that accepts same args as ServeHTTP method also implements http.Handler
