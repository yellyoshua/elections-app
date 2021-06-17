package graphql

import (
	gql "github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
	"github.com/yellyoshua/elections-app/commons/logger"
)

type GraphqlConfig struct {
	Playground bool
	GraphiQL   bool
	Pretty     bool
}

func New(graphql_config GraphqlConfig, graphql_schema_conf gql.SchemaConfig) (*gqlhandler.Handler, *gql.Schema) {
	schema, err := gql.NewSchema(graphql_schema_conf)
	if err != nil {
		logger.Panic("error with graphql schemas -> %s", err)
	}

	return gqlhandler.New(&gqlhandler.Config{
		Schema:     &schema,
		Playground: graphql_config.Playground,
		GraphiQL:   graphql_config.GraphiQL,
		Pretty:     graphql_config.Pretty,
	}), &schema
}

func CreateRootQueries(fields gql.Fields) *gql.Object {
	queries := gql.NewObject(gql.ObjectConfig{
		Name:   "RootQuery",
		Fields: fields,
	})

	return queries
}

func CreateRootMutators(fields gql.Fields) *gql.Object {
	mutators := gql.NewObject(gql.ObjectConfig{
		Name:   "RootMutation",
		Fields: fields,
	})

	return mutators
}

// package graphql

// import (
// 	"net/http"

// 	gql "github.com/graphql-go/graphql"
// 	gqlhandler "github.com/graphql-go/graphql-go-handler"
// 	"github.com/yellyoshua/elections-app/commons/logger"
// 	"github.com/yellyoshua/elections-app/commons/repository"
// )

// // TODO: Create graphql mutator for delete user by username and query find

// // Instance _
// type Instance interface {
// 	Handler() http.Handler
// }

// type instancestruct struct {
// 	schema gql.Schema
// 	config GraphqlConfig
// }

// // GraphqlConfig __
// type GraphqlConfig struct {
// 	Playground bool
// 	GraphiQL   bool
// 	Pretty     bool
// }

// // Initialize func to check if graphql instance funcs correct
// func Initialize() {
// 	repo := repository.New()
// 	_, err := graphqlInit(repo)
// 	if err != nil {
// 		logger.Panic("Error trying setup graphql schemas -> %s", err)
// 	}
// }

// // New return a graphql instance with default repository
// func New(config GraphqlConfig) Instance {
// 	repo := repository.New()
// 	return NewWitRepository(repo, config)
// }

// // NewWitRepository return a graphql instance with a custom repository
// func NewWitRepository(repo repository.Repository, config GraphqlConfig) Instance {
// 	schemas, _ := graphqlInit(repo)
// 	return &instancestruct{schema: schemas, config: config}
// }

// func newGraphqlModulesResolvers(repo repository.Repository) GraphqlModulesResolvers {
// 	return &gqlresolverstruct{
// 		repo: repo,
// 	}
// }

// func graphqlInit(repo repository.Repository) (gql.Schema, error) {
// 	GraphqlResolvers := newGraphqlModulesResolvers(repo)

// 	var mutators = gql.NewObject(gql.ObjectConfig{
// 		Name: "RootMutation",
// 		Fields: gql.Fields{
// 			"updateUser": &gql.Field{
// 				Type:        UserOutputModel, // the return type for this field
// 				Description: "Update a user",
// 				Args:        UpdateUserModel,
// 				Resolve:     GraphqlResolvers.UpdateUser,
// 			},
// 		},
// 	})

// 	var queries = gql.NewObject(gql.ObjectConfig{
// 		Name: "RootQuery",
// 		Fields: gql.Fields{
// 			"users": &gql.Field{
// 				Type:        gql.NewList(UserOutputModel),
// 				Description: "Get users",
// 				Resolve:     GraphqlResolvers.GetUsers,
// 			},
// 			"findUserByID": &gql.Field{
// 				Type:    UserOutputModel,
// 				Args:    FindUserByIDModel,
// 				Resolve: GraphqlResolvers.FindUserByID,
// 			},
// 			"findUserByUsername": &gql.Field{
// 				Type:    UserOutputModel,
// 				Args:    FindUserByUsernameModel,
// 				Resolve: GraphqlResolvers.FindUserByUsername,
// 			},
// 			"createUser": &gql.Field{
// 				Type:    UserOutputModel,
// 				Args:    CreateUserModel,
// 				Resolve: GraphqlResolvers.CreateUser,
// 			},
// 		},
// 	})

// 	var schema, err = gql.NewSchema(gql.SchemaConfig{
// 		Query:    queries,
// 		Mutation: mutators,
// 	})

// 	return schema, err
// }

// // Handler http handler
// func (gqli *instancestruct) Handler() http.Handler {

// 	graphqlHandler := gqlhandler.New(&gqlhandler.Config{
// 		Schema:     &gqli.schema,
// 		Pretty:     gqli.config.Pretty,
// 		Playground: gqli.config.Playground,
// 		GraphiQL:   gqli.config.GraphiQL,
// 	})

// 	return graphqlHandler
// }
