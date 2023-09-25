package routes

import (
	"github.com/johnesleyer/cookieauth/handlers"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/secure", handlers.BasicAuthMiddleware(handlers.SecureHandler))
	r.HandleFunc("/setcookie", handlers.SetCookieHandler)
	r.HandleFunc("/authcookie", handlers.AuthCookieHandler)

	return r
}
