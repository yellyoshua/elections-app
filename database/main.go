package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDatabase connection with database
func ConnectDatabase(mongoURI string, dbName string) *mongo.Database {

	// Set client options
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("err connection database error: ", err)
		return nil
	}

	// Check the connection
	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatal("err ping database error: ", err)
		return nil
	}

	db := client.Database(dbName)
	return db
}
