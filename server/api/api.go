package api

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yellyoshua/elections-app/logger"
	"github.com/yellyoshua/elections-app/server/middlewares"
)

// PublicFolder path for serve static files
var PublicFolder string = "public"

// UploadFolder path for serve static files
var UploadFolder string = "public/uploads"

// NewRestService instance services
func NewRestService(callback func(router *gin.Engine)) {
	var defaultRoute string = ""

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")

	api.Use(middlewares.AuthRequiredMiddleware)
	router.Use(middlewares.CorsMiddleware)

	api.GET(defaultRoute, handlerAPI)
	router.Static("/f/", PublicFolder)
	router.Static("/static/", UploadFolder)
	router.GET("/", handlerHome)

	router.POST("/auth/local", middlewares.BodyLoginUser, handlerLoginUser)
	callback(router)

	server := createServer(router)
	if err := server.ListenAndServe(); err != nil {
		logger.Fatal("Error listen server: %v", err)
	}
}

func createServer(router *gin.Engine) *http.Server {
	var port string = os.Getenv("PORT")

	if len(port) == 0 {
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

func gingonictohttp(handler http.Handler) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		handler.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

// handlerHome rest handler home
func handlerHome(ctx *gin.Context) {
	// r.HeadersRegexp("Content-Type", "application/(text|json)")
	time.Sleep(100 * time.Microsecond)
	ctx.String(http.StatusOK, "Powered with Golang")
}

// HandlerAPI handle api
func handlerAPI(ctx *gin.Context) {
	ctx.String(http.StatusOK, "API - Powered with Golang")
}

// HandlerLoginUser user login post request
func handlerLoginUser(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Powered with Golang")
}
