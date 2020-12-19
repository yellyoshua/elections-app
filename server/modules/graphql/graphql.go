package graphql

import (
	"net/http"

	"github.com/graphql-go/graphql"
	gql "github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
	"github.com/yellyoshua/elections-app/logger"
	"github.com/yellyoshua/elections-app/server/models"
	"github.com/yellyoshua/elections-app/server/services"
)

var gqlSrv services.GRAPHQL

// Initialize func to init graphql module
func Initialize() services.GRAPHQL {
	gqlSrv = services.NewGraphqlService()
	return gqlSrv
}

// Handler http handler
func Handler() http.Handler {

	if gqlSrv == nil {
		logger.ServerFatal("Need initialize graphql module")
	}

	var mutators = graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"updateUser": &graphql.Field{
				Type:        models.UserGQL, // the return type for this field
				Description: "Update a user",
				Args:        models.UpdateUserGQL,
				Resolve:     gqlSrv.UpdateUsers,
			},
		},
	})

	var queries = graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"users": &graphql.Field{
				Type:        graphql.NewList(models.UserGQL),
				Description: "Update a user",
				Resolve:     gqlSrv.GetUsers,
			},
			"findUser": &graphql.Field{
				Type: models.UserGQL,
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: gqlSrv.FindUser,
			},
		},
	})

	var schema, err = gql.NewSchema(gql.SchemaConfig{
		Query:    queries,
		Mutation: mutators,
	})

	// TODO: Test this exeption!
	if err != nil {
		logger.ServerFatal("Error creating graphql schema config, error: %v", err)
	}

	graphqlHandler := gqlhandler.New(&gqlhandler.Config{
		Schema:     &schema,
		Pretty:     true,
		Playground: true,
		GraphiQL:   true,
	})

	return graphqlHandler
}
