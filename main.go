package main

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/johnesleyer/cookieauth/routes"
)

func main() {
	r := routes.NewRouter()

	http.Handle("/", r)

	// Start the server
	http.ListenAndServe(":8080", handlers.LoggingHandler(http.DefaultServeMux, handlers.LogFormatter))
}
