package main

import (
	"net/http"

	"github.com/yellyoshua/elections-app/database"
	"github.com/yellyoshua/elections-app/server"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Database

func main() {
	db = database.ConnectDatabase(GetDatabaseCredentials())

	router := server.CreateServer()
	router.CreateStaticPath("/static/", "public")

	router.Handle("/", http.MethodGet, HandlerHome)
	router.Handle("/auth/local", http.MethodPost, UserLoginValidator(HandlerUserLogin))

	router.HandleGet("/api", HandlerHome, NextMiddleware)

	router.Listen(":3000")

	// server := CreateServer()
	// server.CreateStaticPath("/static/", "public")
	// router := server.Router()
	// router.Handle("/", http.MethodGet, HandlerHome)
	// router.Handle("/auth/local", http.MethodPost, UserLoginValidator(HandlerUserLogin))

	// server.Listen(":3000")
}
