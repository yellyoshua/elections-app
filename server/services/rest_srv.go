package services

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// HandlerHome rest handler home
func (s *Service) HandlerHome(ctx *gin.Context) {
	// r.HeadersRegexp("Content-Type", "application/(text|json)")
	time.Sleep(100 * time.Microsecond)
	ctx.String(http.StatusOK, "Powered with Golang")
}

// HandlerAPI handle api
func (s *Service) HandlerAPI(ctx *gin.Context) {
	ctx.String(http.StatusOK, "API - Powered with Golang")
}

// HandlerLoginUser user login post request
func (s *Service) HandlerLoginUser(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Powered with Golang")
}
