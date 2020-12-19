package modules

import (
	"github.com/yellyoshua/elections-app/logger"
	"github.com/yellyoshua/elections-app/server/modules/graphql"
)

// InitializeModules setup modules environment states
func InitializeModules() {
	graphql.Initialize()

	logger.Server("Modules initialized!")
}
