package gqlresolvers

import (
	gql "github.com/graphql-go/graphql"
	"github.com/yellyoshua/elections-app/server/models"
	"go.mongodb.org/mongo-driver/mongo"
)

// Queries __
type Queries interface {
	GetUsers(params gql.ResolveParams) (interface{}, error)
	FindUser(params gql.ResolveParams) (interface{}, error)
}

// Mutates __
type Mutates interface {
	UpdateUsers(params gql.ResolveParams) (interface{}, error)
}

// GQLMutates __
type GQLMutates struct {
	db *mongo.Database
}

// GQLQueries __
type GQLQueries struct {
	db *mongo.Database
}

// SetupQueries __
func SetupQueries(db *mongo.Database) Queries {
	return &GQLQueries{db: db}
}

// SetupMutates __
func SetupMutates(db *mongo.Database) Mutates {
	return &GQLMutates{db: db}
}

// GetUsers __
func (db *GQLQueries) GetUsers(params gql.ResolveParams) (interface{}, error) {
	return []models.User{}, nil
}

// FindUser __
func (db *GQLQueries) FindUser(params gql.ResolveParams) (interface{}, error) {
	return []models.User{}, nil
}

// UpdateUsers {params models.UpdateUserGQL}
func (db *GQLMutates) UpdateUsers(params gql.ResolveParams) (interface{}, error) {
	username, _ := params.Args["username"].(string)

	return models.User{
		Username: username,
	}, nil
}
