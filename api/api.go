package api

import (
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yellyoshua/elections-app/middlewares"
)

// PublicFolder path for serve static files
var PublicFolder string = "public"

// UploadFolder path for serve static files
var UploadFolder string = "public/uploads"

// API __
type API interface {
	GET(path string, handler Handler)
	POST(path string, handler Handler)
	PUT(path string, handler Handler)
	DELETE(path string, handler Handler)
	Listen(port string) error
}

// Handler __
type Handler func(http.ResponseWriter, *http.Request)

type apistruct struct {
	router *gin.Engine
}

// New instace api service
func New() API {
	var defaultRoute string = ""
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	api := router.Group("/api")

	api.Use(MiddlewareWrapper(middlewares.AuthRequiredMiddleware))
	router.Use(MiddlewareWrapper(middlewares.CorsMiddleware))

	api.GET(defaultRoute, HandlerWrapper(handlerAPI))
	router.Static("/f/", PublicFolder)
	router.Static("/static/", UploadFolder)
	router.GET("/", HandlerWrapper(handlerHome))

	router.POST("/auth/local", MiddlewareWrapper(middlewares.BodyLoginUser), HandlerWrapper(handlerLoginUser))

	return &apistruct{router: router}
}

func (api *apistruct) Listen(port string) error {
	server := createServer(api.router, port)
	return server.ListenAndServe()
}

func (api *apistruct) GET(path string, handler Handler) {
	api.router.GET(path, HandlerWrapper(handler))
}

func (api *apistruct) POST(path string, handler Handler) {
	api.router.POST(path, HandlerWrapper(handler))
}

func (api *apistruct) PUT(path string, handler Handler) {
	api.router.PUT(path, HandlerWrapper(handler))
}

func (api *apistruct) DELETE(path string, handler Handler) {
	api.router.DELETE(path, HandlerWrapper(handler))
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

// MiddlewareWrapper __
func MiddlewareWrapper(handler func(http.ResponseWriter, *http.Request) error) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		if err := handler(ctx.Writer, ctx.Request); err != nil {
			ctx.String(http.StatusUnauthorized, "Unauthorized")
			ctx.AbortWithStatus(http.StatusUnauthorized)
		} else {
			ctx.Next()
		}
	}
}

// HandlerWrapper __
func HandlerWrapper(handler Handler) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		handler(ctx.Writer, ctx.Request)
	}
}

// handlerHome rest handler home
func handlerHome(w http.ResponseWriter, r *http.Request) {
	// r.HeadersRegexp("Content-Type", "application/(text|json)")
	time.Sleep(100 * time.Microsecond)
	ResponseString(w, "Powered with Golang")
}

// HandlerAPI handle api
func handlerAPI(w http.ResponseWriter, r *http.Request) {
	ResponseString(w, "API - Powered with Golang")
}

// HandlerLoginUser user login post request
func handlerLoginUser(w http.ResponseWriter, r *http.Request) {
	ResponseString(w, "Powered with Golang")
}

// ResponseString response a string with status 200
func ResponseString(w http.ResponseWriter, text string) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, text)
}
