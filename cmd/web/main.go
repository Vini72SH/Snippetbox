package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP Network Address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

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

	logger.Info("Starting server", "addr", *addr)

	// Listen port (*addr)
	err := http.ListenAndServe(*addr, mux)

	logger.Error(err.Error())
	os.Exit(1)
}
