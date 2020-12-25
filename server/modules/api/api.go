package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// REST _
type REST interface {
	HandlerHome(ctx *gin.Context)
	HandlerAPI(ctx *gin.Context)
	HandlerLoginUser(ctx *gin.Context)
}

// Service _
type Service int

// NewRestService instance services
func NewRestService() REST {
	service := new(Service)
	return service
}

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
