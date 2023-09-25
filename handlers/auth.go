package handlers

import (
	"fmt"
	"net/http"
)

func SecureHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Authentication using Basic Auth Successful!")
}
