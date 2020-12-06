package middlewares

import (
	"github.com/gin-gonic/gin"
)

// AuthRequiredMiddleware _
func AuthRequiredMiddleware(ctx *gin.Context) {
	ctx.Next()
}

// CorsMiddleware habilitate external request
func CorsMiddleware(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Server", "Powered with Golang")
	ctx.Next()
}
