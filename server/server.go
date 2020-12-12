package server

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
	"github.com/yellyoshua/elections-app/server/database"
)

// ClientDatabase mongo client connection session
func ClientDatabase() *mongo.Database {
	client := database.Connect()
	return client
}

// CreateServer create a server and database connection, this return a gin-gonic router
func CreateServer() *gin.Engine {
	router := gin.Default()
	return router
}
