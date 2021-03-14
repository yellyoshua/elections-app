package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yellyoshua/elections-app/modules/graphql"
)

// HandlerHome rest handler home
func HandlerHome(ctx *gin.Context) {
	// r.HeadersRegexp("Content-Type", "application/(text|json)")
	time.Sleep(100 * time.Microsecond)
	ResponseString(ctx, "Powered with Golang")
}

// HandlerGraphql handler http request of graphql
func HandlerGraphql(ctx *gin.Context) {
	graphqlModule := graphql.Handler()
	graphqlModule.ServeHTTP(ctx.Writer, ctx.Request)
}

// HandlerAPI handle api
func HandlerAPI(ctx *gin.Context) {
	ResponseString(ctx, "API - Powered with Golang")
}

// HandlerLoginUser user login post request
func HandlerLoginUser(ctx *gin.Context) {
	ResponseString(ctx, "Powered with Golang")
}

// ResponseString response a string with status 200
func ResponseString(ctx *gin.Context, text string) {
	ctx.String(http.StatusOK, text)
}
