package setups

import (
	"github.com/yellyoshua/elections-app/server"
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
