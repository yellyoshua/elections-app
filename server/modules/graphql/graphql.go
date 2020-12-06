package graphql

import (
	"fmt"
	"net/http"

	gql "github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
	"github.com/yellyoshua/elections-app/server"
	"github.com/yellyoshua/elections-app/server/database"
)

var queryType = gql.NewObject(gql.ObjectConfig{
	Name: "Query",
	Fields: gql.Fields{
		"latestPost": &gql.Field{
			Type: gql.String,
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				return "Hello World!", nil
			},
		},
		"currentPost": &gql.Field{
			Type: gql.String,
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				return "Hello Post!", nil
			},
		},
	},
})

var Schema, _ = gql.NewSchema(gql.SchemaConfig{
	Query: queryType,
})

// Handler http handler
func Handler() http.Handler {
	fmt.Printf("db: %s", server.DatabaseClient.Collection(database.ProfileCollectionName).Name())
	graphqlHandler := gqlhandler.New(&gqlhandler.Config{
		Schema:     &Schema,
		Pretty:     true,
		Playground: true,
		GraphiQL:   true,
	})

	return graphqlHandler
}
