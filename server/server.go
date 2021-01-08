package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yellyoshua/elections-app/server/modules/api"
	"github.com/yellyoshua/elections-app/server/modules/graphql"
)

// Initialize create a server and database connection, this return a gin-gonic router
func Initialize() {
	HandlerGraphql := graphql.Handler()

	api.NewRestService(func(router *gin.Engine) {
		router.GET("/graphql", gingonictohttp(HandlerGraphql))
		router.POST("/graphql", gingonictohttp(HandlerGraphql))
		router.PUT("/graphql", gingonictohttp(HandlerGraphql))
		router.DELETE("/graphql", gingonictohttp(HandlerGraphql))
		router.OPTIONS("/graphql", gingonictohttp(HandlerGraphql))
	})
}

func gingonictohttp(handler http.Handler) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		handler.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
