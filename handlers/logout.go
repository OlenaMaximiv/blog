package handlers

import (
	"fmt"
	"net/http"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Code to clear authentication session or cookie
	fmt.Fprintf(w, "Logged out successfully")
}
