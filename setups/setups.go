package setups

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/yellyoshua/elections-app/logger"
	"github.com/yellyoshua/elections-app/server"
	"github.com/yellyoshua/elections-app/server/api"
	"github.com/yellyoshua/elections-app/server/modules"
	"github.com/yellyoshua/elections-app/server/repository"
)

// Repositories established connection to database
func Repositories() {
	var indexes bool = true
	repository.Initialize(indexes)
}

// Modules setup modules confs and variables
func Modules() {
	modules.Initialize()
}

// Server start gin-gonic router
func Server() {
	server.Initialize()
}

// Environments if not exist .env file load system environments or defaults!
func Environments() {
	godotenv.Load(".env")

	envs := map[string]bool{
		"PORT":          false,
		"GCS_BUCKET":    true,
		"DATABASE_URI":  true,
		"DATABASE_NAME": true,
	}

	for name, isRequired := range envs {
		if value := os.Getenv(name); len(value) == 0 && isRequired {
			logger.Fatal("%v environment is requerid", name)
		}
	}

	return
}

// Folders create and setup permissions if don't exist
func Folders() {
	folders := map[string]os.FileMode{
		api.PublicFolder: 0755,
		api.UploadFolder: 0755,
	}

	for folder, permission := range folders {
		if notExistFolder(folder) {
			go os.Mkdir(folder, permission)
		}
	}
	return
}

func notExistFolder(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}
