package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc) // registers the handler function and executes that function based on the string pattern
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", nil) // pass nil here it uses the default serve mux as the handler
}
