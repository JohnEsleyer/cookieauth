package handlers

import (
	"fmt"
	"net/http"
)

func SetCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "auth-cookie",
		Value:    "authenticated",
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	fmt.Fprintln(w, "Cookie has been set.")
}

func AuthCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("auth-cookie")
	if err != nil || cookie.Value != "authenticated" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	fmt.Fprintln(w, "Authentication using Cookie Successful!")
}
