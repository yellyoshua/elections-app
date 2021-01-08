package modules

import (
	"github.com/yellyoshua/elections-app/logger"
	"github.com/yellyoshua/elections-app/server/modules/api"
	"github.com/yellyoshua/elections-app/server/modules/graphql"
)

// Initialize create and setup modules
func Initialize() {
	api.Initialize()
	graphql.Initialize()
	logger.Info("Modules initialized")
}
