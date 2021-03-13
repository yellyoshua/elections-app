package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/yellyoshua/elections-app/api"
	"github.com/yellyoshua/elections-app/logger"
	"github.com/yellyoshua/elections-app/modules"
	"github.com/yellyoshua/elections-app/modules/graphql"
	"github.com/yellyoshua/elections-app/repository"
)

func main() {
	setupFolders()
	setupEnvironments()
	setupRepository()
	setupModules()
	setupServer()
}

// Repository established connection to database
func setupRepository() {
	var indexes bool = true
	repository.Initialize(indexes)
}

// Modules setup modules confs and variables
func setupModules() {
	modules.Initialize()
}

// Server start gin-gonic router
func setupServer() {
	HandlerGraphql := graphql.Handler()

	router := api.New()

	router.GET("/graphql", HandlerGraphql.ServeHTTP)
	router.POST("/graphql", HandlerGraphql.ServeHTTP)
	router.PUT("/graphql", HandlerGraphql.ServeHTTP)
	router.DELETE("/graphql", HandlerGraphql.ServeHTTP)

	router.Listen("4000")
}

// Environments if not exist .env file load system environments or defaults!
func setupEnvironments() {
	godotenv.Load(".env")

	envs := map[string]bool{
		"PORT":                 false,
		"GOOGLE_CLOUD_PROJECT": true,
		"GCS_BUCKET":           true,
		"GCS_CREDENTIALS_FILE": true,
		"DATABASE_URI":         true,
		"DATABASE_NAME":        true,
	}

	for name, isRequired := range envs {
		value := os.Getenv(name)

		if existEnv := len(value) == 0; existEnv && isRequired {
			logger.Fatal("%v environment variable must be set.\n", name)
		}
	}

	return
}

// Folders create and setup permissions if don't exist
func setupFolders() {
	folders := map[string]os.FileMode{
		api.PublicFolder: os.ModeDir,
		api.UploadFolder: os.ModeDir,
	}

	for folder, permission := range folders {
		if notExistFolder := checkNoFolder(folder); notExistFolder {
			go os.Mkdir(folder, permission)
		}
	}
	return
}

func checkNoFolder(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}
