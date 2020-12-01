package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// UserCollectionName collection name in db
const (
	UserCollectionName    = "users"
	SessionCollectionName = "sessions"
	ProfileCollectionName = "profiles"
)

// GetDatabaseCredentials load .env file and extract vars
func GetDatabaseCredentials() (string, string) {
	godotenv.Load(".env")

	var (
		dbName   string = os.Getenv("DB_NAME")
		mongoURI string = fmt.Sprintf("mongodb://%s:%s@%s:%s/",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_ADDR"),
			os.Getenv("DB_PORT"),
		)
	)

	return mongoURI, dbName
}
