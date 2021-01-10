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
	FindUserByID(params gql.ResolveParams) (interface{}, error)
	FindUserByUsername(params gql.ResolveParams) (interface{}, error)
	CreateUser(params gql.ResolveParams) (interface{}, error)
	// Mutators
	UpdateUser(params gql.ResolveParams) (interface{}, error)
}

// Initialize func to init graphql module
func Initialize() GRAPHQL {
	graphqlService = NewGraphqlService()
	return graphqlService
}

// Service _
type Service struct{}

// NewGraphqlService intance service
func NewGraphqlService() *Service {
	service := new(Service)
	return service
}

// Handler http handler
func Handler() http.Handler {

	if graphqlService == nil {
		logger.Fatal("Not initialized graphql module")
	}

	schema, err := setupSchemas(graphqlService)

	// TODO: Test this exeption!
	if err != nil {
		logger.Fatal("Error creating graphql schema config, error: %v", err)
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
				Resolve:     graphqlService.UpdateUser,
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
			"findUserByID": &graphql.Field{
				Type:    models.UserGQL,
				Args:    models.FindUserByIDGQL,
				Resolve: graphqlService.FindUserByID,
			},
			"findUserByUsername": &graphql.Field{
				Type:    models.UserGQL,
				Args:    models.FindUserByUsernameGQL,
				Resolve: graphqlService.FindUserByUsername,
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

// FindUserByID __
func (s *Service) FindUserByID(params gql.ResolveParams) (interface{}, error) {
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

// FindUserByUsername __
func (s *Service) FindUserByUsername(params gql.ResolveParams) (interface{}, error) {

	col := repository.NewRepository(repository.CollectionUsers)
	username, _ := params.Args["username"].(string)

	var user models.User
	err := col.FindOne(bson.M{"username": username}, &user)

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

	user := models.User{
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
	if err != nil {
		return nil, err
	}

	user.ID = id
	return user, err
}

// UpdateUser {params models.UpdateUserGQL}
func (s *Service) UpdateUser(params gql.ResolveParams) (interface{}, error) {
	var err error
	var userID primitive.ObjectID

	col := repository.NewRepository(repository.CollectionUsers)
	userIDString, _ := params.Args["userID"].(string)
	delete(params.Args, "userID")

	userID, err = primitive.ObjectIDFromHex(userIDString)
	if err != nil {
		return nil, err
	}

	err = col.UpdateOne(bson.M{"_id": userID}, params.Args)

	if err != nil {
		return nil, err
	}

	params.Args["_id"] = userID
	return params.Args, nil
}

func checkError(err error) {
	if err != nil {
		logger.Fatal("Repository error: %v", err)
	}
}
