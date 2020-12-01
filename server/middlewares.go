package server

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware log register every request
func LoggingMiddleware(next http.Handler) http.Handler {
	start := time.Now()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		defer log.Printf(" [%s] %s %s", r.Method, r.RequestURI, duration)
	})
}

// CorsMiddleware habilitate external request
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Server", "Powered with Golang")
		next.ServeHTTP(w, r)
	})
}
