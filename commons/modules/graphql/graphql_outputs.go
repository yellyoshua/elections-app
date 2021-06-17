package graphql

import gql "github.com/graphql-go/graphql"

// UserOutputModel _
var UserOutputModel = gql.NewObject(gql.ObjectConfig{
	Name: "User",
	Fields: gql.Fields{
		"_id": &gql.Field{
			Type: gql.String,
		},
		"name": &gql.Field{
			Type: gql.String,
		},
		"surname": &gql.Field{
			Type: gql.String,
		},
		"fullName": &gql.Field{
			Type: gql.String,
		},
		"username": &gql.Field{
			Type: gql.String,
		},
		"email": &gql.Field{
			Type: gql.String,
		},
		"verified": &gql.Field{
			Type: gql.Boolean,
		},
		"active": &gql.Field{
			Type: gql.Boolean,
		},
		"created": &gql.Field{
			Type: gql.Int,
		},
	},
})
