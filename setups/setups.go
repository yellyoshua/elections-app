package setups

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/yellyoshua/elections-app/server"
	"github.com/yellyoshua/elections-app/server/handlers"
	"github.com/yellyoshua/elections-app/server/middlewares"
	"github.com/yellyoshua/elections-app/server/modules/graphql"
	"github.com/yellyoshua/elections-app/server/validators"
	"go.mongodb.org/mongo-driver/mongo"
)

var clientDatabase *mongo.Database

// Database connect to database
func Database() {
	clientDatabase = server.ClientDatabase()
}

// Server start gin-gonic router
func Server() {
	var port string = os.Getenv("PORT")

	if clientDatabase == nil {
		log.Fatal("step setup database skipped")
	}

	HandlerGraphql := graphql.Handler(clientDatabase)
	router := server.CreateServer()
	router.Use(middlewares.CorsMiddleware)

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
