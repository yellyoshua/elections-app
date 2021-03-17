package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/yellyoshua/elections-app/api"
	"github.com/yellyoshua/elections-app/constants"
	"github.com/yellyoshua/elections-app/handlers"
	"github.com/yellyoshua/elections-app/logger"
	"github.com/yellyoshua/elections-app/middlewares"
	"github.com/yellyoshua/elections-app/modules"
	"github.com/yellyoshua/elections-app/modules/graphql"
	"github.com/yellyoshua/elections-app/repository"
)

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
	port := os.Getenv("PORT")
	router := api.New()

	graphqlSession := graphql.New(graphql.GraphqlConfig{Playground: true, GraphiQL: true, Pretty: true})
	middlw := middlewares.New(middlewares.MiddlewareConf{})

	graphqlHandler := graphqlSession.Handler()

	router.POST("/graphql", api.WrapperGinHandler(graphqlHandler))
	router.GET("/graphql", api.WrapperGinHandler(graphqlHandler))
	router.PUT("/graphql", api.WrapperGinHandler(graphqlHandler))
	router.DELETE("/graphql", api.WrapperGinHandler(graphqlHandler))
	router.Use(middlw.AuthRequiredMiddleware).GET("/api", handlers.HandlerAPI)
	router.Use(middlw.BodyLoginUser).POST("/auth/local", handlers.HandlerLoginUser)
	router.GET("/", handlers.HandlerHome)
	router.Listen(port)
}

// Environments if not exist .env file load system environments or defaults!
func setupEnvironments() {
	godotenv.Load(".env")

	envs := map[string]bool{
		"PORT":                 false,
		"GOOGLE_CLOUD_PROJECT": true,
		"GCS_BUCKET":           true,
		"GCS_CREDENTIALS_FILE": true,
		"MONGODB_URI":          true,
		"MONGODB_DATABASE":     false,
	}

	for name, isRequired := range envs {
		value := os.Getenv(name)

		if existEnv := len(value) == 0; existEnv && isRequired {
			logger.Fatal("%v environment variable is required.\n", name)
		}
	}

	return
}

// Folders create and setup permissions if don't exist
func setupFolders() {
	folders := map[string]os.FileMode{
		constants.APIPublicFolder: os.ModeDir,
		constants.APIUploadFolder: os.ModeDir,
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

func main() {
	setupFolders()
	setupEnvironments()
	setupRepository()
	setupModules()
	setupServer()
}
