package main

import (
	"net/http"
)

// NextMiddleware ignore middleware and next handle
func NextMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(handler.ServeHTTP)
}
