package modules

import (
	"github.com/yellyoshua/elections-app/logger"
	"github.com/yellyoshua/elections-app/modules/authentication"
	"github.com/yellyoshua/elections-app/modules/graphql"
	"github.com/yellyoshua/elections-app/modules/storage"
)

// Initialize create and setup modules
func Initialize() {
	storage.Initialize()
	authentication.Initialize()
	graphql.Initialize()
}

func checkError(err error) {
	if err != nil {
		logger.Fatal("%v", err)
	}
	return
}
