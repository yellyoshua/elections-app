package server

import (
	"github.com/gin-gonic/gin"
	"github.com/yellyoshua/elections-app/logger"
)

// CreateServer create a server and database connection, this return a gin-gonic router
func CreateServer() *gin.Engine {
	logger.Server("Starting GIN-GONIC server")
	router := gin.Default()
	return router
}
