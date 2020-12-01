package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// CreateServer create and return a router with load default middleware
func CreateServer() Router {
	router := mux.NewRouter()
	router.Use(CorsMiddleware)
	router.Use(LoggingMiddleware)

	return &HTTP{
		router: router,
	}
}

// Handle method of router
func (s *HTTP) Handle(path string, method string, httpFun http.HandlerFunc) {
	s.router.HandleFunc(path, httpFun).Methods(method)
}

// HandleGet method get of router
func (s *HTTP) HandleGet(path string, httpFun http.HandlerFunc, middleware Middleware) {
	s.router.HandleFunc(path, middleware(httpFun)).Methods(http.MethodGet)
}

// CreateStaticPath serve files under http://localhost/{path}/<filename>
func (s *HTTP) CreateStaticPath(path string, dir http.Dir) {
	s.router.PathPrefix(path).Handler(http.StripPrefix(path, http.FileServer(http.Dir(dir))))
}

// Listen func listen server with port arg provider
func (s *HTTP) Listen(port string) error {
	server := &http.Server{
		Handler: s.router,
		Addr:    port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("unable run server ", err)
		return err
	}
	log.Print("server running on port ", port)
	return nil
}
