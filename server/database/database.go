package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect connection with database
func Connect() *mongo.Database {
	var (
		dbName   string = os.Getenv("DB_NAME")
		mongoURI string = fmt.Sprintf("mongodb://%s:%s@%s:%s/",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_ADDR"),
			os.Getenv("DB_PORT"),
		)
	)

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
