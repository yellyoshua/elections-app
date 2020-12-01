package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HandlerHome handle home
func HandlerHome(w http.ResponseWriter, r *http.Request) {
	// r.HeadersRegexp("Content-Type", "application/(text|json)")

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Powered with Golang")
}

// HandlerUserLogin user login post request
func HandlerUserLogin(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Powered with Golang")
}
