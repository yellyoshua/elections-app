package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// HandlerHome handle home
func HandlerHome(ctx *gin.Context) {
	// r.HeadersRegexp("Content-Type", "application/(text|json)")
	time.Sleep(100 * time.Microsecond)
	ctx.String(http.StatusOK, "Powered with Golang")
}

// HandlerAPI handle api
func HandlerAPI(ctx *gin.Context) {
	ctx.String(http.StatusOK, "API - Powered with Golang")
}

// HandlerUserLogin user login post request
func HandlerUserLogin(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Powered with Golang")
}
