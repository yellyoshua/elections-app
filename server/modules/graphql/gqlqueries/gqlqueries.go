package gqlqueries

import (
	"github.com/graphql-go/graphql"
	"github.com/yellyoshua/elections-app/server/models"
	"github.com/yellyoshua/elections-app/server/modules/graphql/gqlresolvers"
	"go.mongodb.org/mongo-driver/mongo"
)

// Setup __
func Setup(db *mongo.Database) *graphql.Object {
	queries := gqlresolvers.SetupQueries(db)

	var rootQuery = graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"users": &graphql.Field{
				Type:        graphql.NewList(models.UserGQL),
				Description: "Update a user",
				Resolve:     queries.GetUsers,
			},
			"findUser": &graphql.Field{
				Type: models.UserGQL,
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: queries.FindUser,
			},
		},
	})
	return rootQuery
}
