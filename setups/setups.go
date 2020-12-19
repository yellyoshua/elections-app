package setups

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/yellyoshua/elections-app/server"
	"github.com/yellyoshua/elections-app/server/middlewares"
	"github.com/yellyoshua/elections-app/server/modules"
	"github.com/yellyoshua/elections-app/server/modules/graphql"
	"github.com/yellyoshua/elections-app/server/repository"
	"github.com/yellyoshua/elections-app/server/services"
)

// Repositories established connection to database
func Repositories() {
	repository.Initialize()
}

// Modules setup modules confs and variables
func Modules() {
	modules.InitializeModules()
}

// Server start gin-gonic router
func Server() {

	var port string = os.Getenv("PORT")

	restSrv := services.NewRestService()

	HandlerGraphql := graphql.Handler()
	router := server.CreateServer()
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
