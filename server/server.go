package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yellyoshua/elections-app/logger"
	"github.com/yellyoshua/elections-app/server/middlewares"
	"github.com/yellyoshua/elections-app/server/modules/api"
	"github.com/yellyoshua/elections-app/server/modules/graphql"
)

var router *gin.Engine

// PublicFolder path for serve static files
var PublicFolder string = "public"

// UploadFolder path for serve static files
var UploadFolder string = "public/uploads"

// Initialize create a server and database connection, this return a gin-gonic router
func Initialize(port string) *gin.Engine {
	logger.Info("Starting GIN-GONIC server")
	router = gin.Default()
	start(port)
	return router
}

func start(port string) {
	restSrv := api.NewRestService()
	HandlerGraphql := graphql.Handler()

	router.Use(middlewares.CorsMiddleware)

	router.Static("/f/", PublicFolder)
	router.Static("/static/", UploadFolder)

	router.GET("/graphql", gingonictohttp(HandlerGraphql))
	router.POST("/graphql", gingonictohttp(HandlerGraphql))
	router.PUT("/graphql", gingonictohttp(HandlerGraphql))
	router.DELETE("/graphql", gingonictohttp(HandlerGraphql))
	router.OPTIONS("/graphql", gingonictohttp(HandlerGraphql))

	router.GET("/", restSrv.HandlerHome)
	router.GET("/api", restSrv.HandlerAPI)
	router.POST("/auth/local", middlewares.BodyLoginUser, restSrv.HandlerLoginUser)

	router.Run(":" + port)
}

func gingonictohttp(handler http.Handler) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		handler.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
