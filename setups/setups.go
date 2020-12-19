package setups

import (
	"os"

	"github.com/yellyoshua/elections-app/server"
	"github.com/yellyoshua/elections-app/server/modules"
	"github.com/yellyoshua/elections-app/server/repository"
)

// Repositories established connection to database
func Repositories() {
	repository.Initialize()
}

// Modules setup modules confs and variables
func Modules() {
	modules.Initialize()
}

// Server start gin-gonic router
func Server() {
	var port string = os.Getenv("PORT")
	server.Initialize(port)
}
