package graphql

import gql "github.com/graphql-go/graphql"

// CreateUserModel args graphql create user query
var CreateUserModel = gql.FieldConfigArgument{
	"name": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"surname": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"username": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"email": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"password": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
}

// FindUserByIDModel args graphql
var FindUserByIDModel = gql.FieldConfigArgument{
	"id": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
}

// FindUserByUsernameModel args graphql
var FindUserByUsernameModel = gql.FieldConfigArgument{
	"username": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
}

// UpdateUserModel args graphql update user query
var UpdateUserModel = gql.FieldConfigArgument{
	"userID": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"name": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"surname": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"fullName": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"username": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"email": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"verified": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.Boolean),
	},
	"active": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.Boolean),
	},
}
