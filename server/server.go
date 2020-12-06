package server

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
	"github.com/yellyoshua/elections-app/server/database"
)

// DatabaseClient mongo client connection session
var DatabaseClient *mongo.Database

// CreateServer create a server and database connection, this return a gin-gonic router
func CreateServer(testing bool) *gin.Engine {
	if testing != true {
		DatabaseClient = database.Connect()
	}
	router := gin.Default()
	return router
}
