package modules

import (
	logger "github.com/yellyoshua/elections-app/logger"
	authentication "github.com/yellyoshua/elections-app/modules/authentication"
	graphql "github.com/yellyoshua/elections-app/modules/graphql"
	storage "github.com/yellyoshua/elections-app/modules/storage"
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
