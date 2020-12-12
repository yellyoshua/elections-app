package gqlmutators

import (
	"github.com/graphql-go/graphql"
	"github.com/yellyoshua/elections-app/server/models"
	"github.com/yellyoshua/elections-app/server/modules/graphql/gqlresolvers"
	"go.mongodb.org/mongo-driver/mongo"
)

// Setup __
func Setup(db *mongo.Database) *graphql.Object {
	mutators := gqlresolvers.SetupMutates(db)

	var rootMutation = graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"updateUser": &graphql.Field{
				Type:        models.UserGQL, // the return type for this field
				Description: "Update a user",
				Args:        models.UpdateUserGQL,
				Resolve:     mutators.UpdateUsers,
			},
		},
	})

	return rootMutation
}
