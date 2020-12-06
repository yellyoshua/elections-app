package setups

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/yellyoshua/elections-app/server"
	"github.com/yellyoshua/elections-app/server/handlers"
	"github.com/yellyoshua/elections-app/server/middlewares"
	"github.com/yellyoshua/elections-app/server/modules/graphql"
	"github.com/yellyoshua/elections-app/server/validators"
)

// ServerAndDatabase connect to database and start gin-gonic router
func ServerAndDatabase(isTesting bool) {
	var port string = os.Getenv("PORT")
	router := server.CreateServer(isTesting)
	router.Use(middlewares.CorsMiddleware)

	HandlerGraphql := graphql.Handler()

	router.Static("/f/", PublicFolder)
	router.Static("/static/", UploadFolder)

	router.GET("/graphql", gingonictohttp(HandlerGraphql))
	router.POST("/graphql", gingonictohttp(HandlerGraphql))
	router.PUT("/graphql", gingonictohttp(HandlerGraphql))
	router.DELETE("/graphql", gingonictohttp(HandlerGraphql))
	router.OPTIONS("/graphql", gingonictohttp(HandlerGraphql))

	router.GET("/", handlers.HandlerHome)
	router.GET("/api", handlers.HandlerAPI)
	router.POST("/auth/local", validators.UserLoginValidator, handlers.HandlerUserLogin)

	router.Run(":" + port)
}

func gingonictohttp(handler http.Handler) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		handler.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
