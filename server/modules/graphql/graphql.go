package graphql

import (
	"net/http"

	gql "github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
	"github.com/yellyoshua/elections-app/server/modules/graphql/gqlmutators"
	"github.com/yellyoshua/elections-app/server/modules/graphql/gqlqueries"
	"go.mongodb.org/mongo-driver/mongo"
)

// Handler http handler
func Handler(db *mongo.Database) http.Handler {
	var queries = gqlqueries.Setup(db)
	var mutators = gqlmutators.Setup(db)

	var schema, _ = gql.NewSchema(gql.SchemaConfig{
		Query:    queries,
		Mutation: mutators,
	})
	// Here resolve schemes/queries/resolvers

	graphqlHandler := gqlhandler.New(&gqlhandler.Config{
		Schema:     &schema,
		Pretty:     true,
		Playground: true,
		GraphiQL:   true,
	})

	return graphqlHandler
}
