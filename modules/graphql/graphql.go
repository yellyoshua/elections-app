package graphql

import (
	"fmt"
	"net/http"
	"time"

	"github.com/graphql-go/graphql"
	gql "github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
	"github.com/yellyoshua/elections-app/logger"
	"github.com/yellyoshua/elections-app/models"
	"github.com/yellyoshua/elections-app/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO: Create graphql mutator for delete user by username and query find

// resolver Queries and Mutators
type resolver interface {
	// Queries
	GetUsers(params gql.ResolveParams) (interface{}, error)
	FindUserByID(params gql.ResolveParams) (interface{}, error)
	FindUserByUsername(params gql.ResolveParams) (interface{}, error)
	CreateUser(params gql.ResolveParams) (interface{}, error)
	// Mutators
	UpdateUser(params gql.ResolveParams) (interface{}, error)
}

// Service _
type Service interface {
	Handler() http.Handler
}

type serviceStruct struct {
	schema gql.Schema
}

type resolverStruct struct {
	repo repository.Repository
}

// Initialize func to init graphql module
func Initialize() {
	repo := repository.New()
	_, err := graphqlSchemas(repo)
	if err != nil {
		logger.Fatal("Error trying setup graphql schemas -> %s", err)
	}
}

// New _
func New() Service {
	repo := repository.New()
	return NewWitRepository(repo)
}

// NewWitRepository _
func NewWitRepository(repo repository.Repository) Service {
	schema, _ := graphqlSchemas(repo)
	return &serviceStruct{schema: schema}
}

func graphqlSchemas(repo repository.Repository) (gql.Schema, error) {
	graphqlService := &resolverStruct{
		repo: repo,
	}

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

// Handler http handler
func (gql *serviceStruct) Handler() http.Handler {

	graphqlHandler := gqlhandler.New(&gqlhandler.Config{
		Schema:     &gql.schema,
		Pretty:     true,
		Playground: true,
		GraphiQL:   true,
	})

	return graphqlHandler
}

func (r *resolverStruct) GetUsers(params gql.ResolveParams) (interface{}, error) {
	col := r.repo.Col(repository.CollectionUsers)

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

	col := r.repo.Col(repository.CollectionUsers)
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

	col := r.repo.Col(repository.CollectionUsers)
	username, _ := params.Args["username"].(string)

	var user models.User
	err := col.FindOne(bson.M{"username": username}, &user)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *resolverStruct) CreateUser(params gql.ResolveParams) (interface{}, error) {
	col := r.repo.Col(repository.CollectionUsers)
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

	col := r.repo.Col(repository.CollectionUsers)
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
