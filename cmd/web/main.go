package main

import (
	"log"
	"net/http"
)

func main() {

	// Start web server
	mux := http.NewServeMux()

	// Create a file server with the files in ./ui/static/
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Defines the GET method to obtain static files at runtime
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

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
