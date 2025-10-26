package main

import (
	"log"
	"net/http"
)

func main() {

	// Start web server
	mux := http.NewServeMux()

	// Handlers and Routes
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("Starting sever on :4000")

	// Listen port 4000
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
