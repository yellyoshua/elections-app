package graphql

import (
	"fmt"
	"net/http"
	"time"

	gql "github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
	"github.com/yellyoshua/elections-app/constants"
	"github.com/yellyoshua/elections-app/logger"
	"github.com/yellyoshua/elections-app/models"
	"github.com/yellyoshua/elections-app/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO: Create graphql mutator for delete user by username and query find

// grqphqlresolver Queries and Mutators
type grqphqlresolver interface {
	// Queries
	GetUsers(params gql.ResolveParams) (interface{}, error)
	FindUserByID(params gql.ResolveParams) (interface{}, error)
	FindUserByUsername(params gql.ResolveParams) (interface{}, error)
	CreateUser(params gql.ResolveParams) (interface{}, error)
	// Mutators
	UpdateUser(params gql.ResolveParams) (interface{}, error)
}

type resolverStruct struct {
	repo repository.Repository
}

// Service _
type Service interface {
	Handler() http.Handler
}

type serviceStruct struct {
	schema gql.Schema
	config GraphqlConfig
}

// GraphqlConfig __
type GraphqlConfig struct {
	Playground bool
	GraphiQL   bool
	Pretty     bool
}

// Initialize func to init graphql module
func Initialize() {
	repo := repository.New()
	_, err := graphqlInit(repo)
	if err != nil {
		logger.Fatal("Error trying setup graphql schemas -> %s", err)
	}
}

// New _
func New(config GraphqlConfig) Service {
	repo := repository.New()
	return NewWitRepository(repo, config)
}

// NewWitRepository _
func NewWitRepository(repo repository.Repository, config GraphqlConfig) Service {
	schemas, _ := graphqlInit(repo)
	return &serviceStruct{schema: schemas, config: config}
}

func newGraphqlResolvers(repo repository.Repository) grqphqlresolver {
	return &resolverStruct{
		repo: repo,
	}
}

func graphqlInit(repo repository.Repository) (gql.Schema, error) {
	graphqlService := newGraphqlResolvers(repo)

	var mutators = gql.NewObject(gql.ObjectConfig{
		Name: "RootMutation",
		Fields: gql.Fields{
			"updateUser": &gql.Field{
				Type:        models.GraphqlUser, // the return type for this field
				Description: "Update a user",
				Args:        models.GraphqlUpdateUser,
				Resolve:     graphqlService.UpdateUser,
			},
		},
	})

	var queries = gql.NewObject(gql.ObjectConfig{
		Name: "RootQuery",
		Fields: gql.Fields{
			"users": &gql.Field{
				Type:        gql.NewList(models.GraphqlUser),
				Description: "Get users",
				Resolve:     graphqlService.GetUsers,
			},
			"findUserByID": &gql.Field{
				Type:    models.GraphqlUser,
				Args:    models.GraphqlFindUserByID,
				Resolve: graphqlService.FindUserByID,
			},
			"findUserByUsername": &gql.Field{
				Type:    models.GraphqlUser,
				Args:    models.GraphqlFindUserByUsername,
				Resolve: graphqlService.FindUserByUsername,
			},
			"createUser": &gql.Field{
				Type:    models.GraphqlUser,
				Args:    models.GraphqlCreateUser,
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

// Handler http handler
func (gql *serviceStruct) Handler() http.Handler {

	graphqlHandler := gqlhandler.New(&gqlhandler.Config{
		Schema:     &gql.schema,
		Pretty:     gql.config.Pretty,
		Playground: gql.config.Playground,
		GraphiQL:   gql.config.GraphiQL,
	})

	return graphqlHandler
}

func (r *resolverStruct) GetUsers(params gql.ResolveParams) (interface{}, error) {
	col := r.repo.Col(constants.CollectionUsers)

	filter := bson.D{}
	var users []models.User
	err := col.Find(filter, &users)

	if err != nil {
		return make([]models.User, 0), err
	}
	return users, nil
}

func (r *resolverStruct) FindUserByID(params gql.ResolveParams) (interface{}, error) {
	var err error
	var userID primitive.ObjectID

	col := r.repo.Col(constants.CollectionUsers)
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

func (r *resolverStruct) FindUserByUsername(params gql.ResolveParams) (interface{}, error) {

	col := r.repo.Col(constants.CollectionUsers)
	username, _ := params.Args["username"].(string)

	var user models.User
	err := col.FindOne(bson.M{"username": username}, &user)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *resolverStruct) CreateUser(params gql.ResolveParams) (interface{}, error) {
	col := r.repo.Col(constants.CollectionUsers)
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

func (r *resolverStruct) UpdateUser(params gql.ResolveParams) (interface{}, error) {
	var err error
	var userID primitive.ObjectID

	col := r.repo.Col(constants.CollectionUsers)
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
