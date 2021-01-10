package modules

import (
	"github.com/yellyoshua/elections-app/server/modules/authentication"
	"github.com/yellyoshua/elections-app/server/modules/graphql"
)

// Initialize create and setup modules
func Initialize() {
	authentication.Initialize()
	graphql.Initialize()
}
