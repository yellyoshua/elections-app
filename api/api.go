package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yellyoshua/elections-app/logger"
)

// PublicFolder path for serve static files
var PublicFolder string = "public"

// UploadFolder path for serve static files
var UploadFolder string = "public/uploads"

// API __
type API interface {
	GET(path string, handler ...gin.HandlerFunc)
	POST(path string, handler ...gin.HandlerFunc)
	PUT(path string, handler ...gin.HandlerFunc)
	DELETE(path string, handler ...gin.HandlerFunc)
	Serve(w http.ResponseWriter, req *http.Request)
	Use(middlewares ...gin.HandlerFunc) API
	Listen(port string) error
}

// Handler __
type Handler gin.HandlerFunc

type apistruct struct {
	router      *gin.Engine
	middlewares []gin.HandlerFunc
}

// New instace api service
func New() API {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Static("/f/", PublicFolder)
	router.Static("/static/", UploadFolder)

	return &apistruct{router: router, middlewares: make([]gin.HandlerFunc, 0)}
}

func (api *apistruct) Use(middlewares ...gin.HandlerFunc) API {
	api.middlewares = middlewares
	return api
}

func (api *apistruct) Serve(w http.ResponseWriter, req *http.Request) {
	middlewares := api.middlewares
	api.middlewares = make([]gin.HandlerFunc, 0)
	api.router.Use(middlewares...)
	api.router.ServeHTTP(w, req)
}

func (api *apistruct) Listen(port string) error {
	server := createServer(api.router, port)
	return server.ListenAndServe()
}

func (api *apistruct) GET(path string, handler ...gin.HandlerFunc) {
	middlewares := api.middlewares
	api.middlewares = make([]gin.HandlerFunc, 0)
	logger.Info("c (%v) o (%v)", len(middlewares), len(api.middlewares))
	api.router.GET(path, append(middlewares, handler...)...)
}

func (api *apistruct) POST(path string, handler ...gin.HandlerFunc) {
	middlewares := api.middlewares
	api.middlewares = make([]gin.HandlerFunc, 0)
	api.router.POST(path, append(middlewares, handler...)...)
}

func (api *apistruct) PUT(path string, handler ...gin.HandlerFunc) {
	middlewares := api.middlewares
	api.middlewares = make([]gin.HandlerFunc, 0)
	api.router.PUT(path, append(middlewares, handler...)...)
}

func (api *apistruct) DELETE(path string, handler ...gin.HandlerFunc) {
	middlewares := api.middlewares
	api.middlewares = make([]gin.HandlerFunc, 0)
	api.router.DELETE(path, append(middlewares, handler...)...)
}

func createServer(router *gin.Engine, port string) *http.Server {
	if noPort := len(port) == 0; noPort {
		port = "3000"
	}

	server := new(http.Server)
	server.Addr = ":" + port
	server.Handler = router
	server.ReadTimeout = 10 * time.Second
	server.WriteTimeout = 10 * time.Second
	server.MaxHeaderBytes = 1 << 20
	return server
}

// WrapperGinHandler pass a parameter a http handler that combine with gin-gonic handler
func WrapperGinHandler(handler http.Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		handler.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
