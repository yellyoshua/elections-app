package graphql

import (
	"fmt"
	"net/http"
	"time"

	"github.com/graphql-go/graphql"
	gql "github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
	"github.com/yellyoshua/elections-app/logger"
	"github.com/yellyoshua/elections-app/server/models"
	"github.com/yellyoshua/elections-app/server/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO: Create graphql mutator for delete user by username and query find

var graphqlService GRAPHQL

// GRAPHQL Queries and Mutators
type GRAPHQL interface {
	// Queries
	GetUsers(params gql.ResolveParams) (interface{}, error)
	FindUser(params gql.ResolveParams) (interface{}, error)
	CreateUser(params gql.ResolveParams) (interface{}, error)
	// Mutators
	UpdateUsers(params gql.ResolveParams) (interface{}, error)
}

// Initialize func to init graphql module
func Initialize() GRAPHQL {
	graphqlService = NewGraphqlService()
	return graphqlService
}

// Service _
type Service int

// NewGraphqlService intance service
func NewGraphqlService() GRAPHQL {
	service := new(Service)
	return service
}

// Handler http handler
func Handler() http.Handler {

	if graphqlService == nil {
		logger.ServerFatal("Need initialize graphql module")
	}

	schema, err := setupSchemas(graphqlService)

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

func setupSchemas(graphqlService GRAPHQL) (gql.Schema, error) {
	var mutators = graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"updateUser": &graphql.Field{
				Type:        models.UserGQL, // the return type for this field
				Description: "Update a user",
				Args:        models.UpdateUserGQL,
				Resolve:     graphqlService.UpdateUsers,
			},
		},
	})

	var queries = graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"users": &graphql.Field{
				Type:        graphql.NewList(models.UserGQL),
				Description: "Get users",
				Resolve:     graphqlService.GetUsers,
			},
			"findUser": &graphql.Field{
				Type:    models.UserGQL,
				Args:    models.FindUser,
				Resolve: graphqlService.FindUser,
			},
			"createUser": &graphql.Field{
				Type:    models.UserGQL,
				Args:    models.CreateUserGQL,
				Resolve: graphqlService.CreateUser,
			},
		},
	})

	var schema, err = gql.NewSchema(gql.SchemaConfig{
		Query:    queries,
		Mutation: mutators,
	})
	return schema, err
}

// GetUsers __
func (s *Service) GetUsers(params gql.ResolveParams) (interface{}, error) {
	col := repository.NewRepository(repository.CollectionUsers)
	filter := bson.D{}
	var users []models.User
	err := col.Find(filter, &users)

	if err != nil {
		return make([]models.User, 0), err
	}
	return users, nil
}

// FindUser __
func (s *Service) FindUser(params gql.ResolveParams) (interface{}, error) {
	var err error
	var userID primitive.ObjectID

	col := repository.NewRepository(repository.CollectionUsers)
	userIDString, _ := params.Args["id"].(string)

	userID, err = primitive.ObjectIDFromHex(userIDString)
	if err != nil {
		return nil, err
	}

	var user models.User
	err = col.FindByID(userID, &user)

	if err != nil {
		return nil, err
	}
	return user, nil
}

// CreateUser __
func (s *Service) CreateUser(params gql.ResolveParams) (interface{}, error) {
	col := repository.NewRepository(repository.CollectionUsers)
	name, _ := params.Args["name"].(string)
	surname, _ := params.Args["surname"].(string)
	username, _ := params.Args["username"].(string)
	email, _ := params.Args["email"].(string)
	password, _ := params.Args["password"].(string)

	user := &models.User{
		Name:     name,
		Surname:  surname,
		Username: username,
		Fullname: fmt.Sprintf("%s %s", name, surname),
		Email:    email,
		Password: password,
		Active:   true,
		Verified: false,
		Created:  time.Now().Unix(),
	}

	id, err := col.InsertOne(user)
	user.ID = id
	return user, err
}

// UpdateUsers {params models.UpdateUserGQL}
func (s *Service) UpdateUsers(params gql.ResolveParams) (interface{}, error) {
	username, _ := params.Args["username"].(string)

	return models.User{
		Username: username,
	}, nil
}

func checkError(err error) {
	if err != nil {
		logger.ServerFatal("repository error: %v", err)
	}
}
