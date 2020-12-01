package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

// HTTP struct server with mux router
type HTTP struct {
	router *mux.Router
}

// Middleware _
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Router _
type Router interface {
	Handle(path string, method string, httpFun http.HandlerFunc)
	HandleGet(path string, httpFun http.HandlerFunc, middleware Middleware)
	CreateStaticPath(path string, dir http.Dir)
	Listen(port string) error
}
