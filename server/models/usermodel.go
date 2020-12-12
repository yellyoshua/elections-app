package models

import (
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User user
type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Username  string             `bson:"username" json:"username"`
	Name      string             `bson:"name" json:"name"`
	Surname   string             `bson:"surname" json:"surname"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password"`
	Validated bool               `bson:"validated" json:"validated,omitempty"`
	Birthday  int64              `bson:"birthday" json:"birthday,omitempty"`
	// Created   int64              `bson:"created" json:"created"`
	Profile primitive.ObjectID `bson:"profile" json:"profile"`
}

// UserGQL _
var UserGQL = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"surname": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"validated": &graphql.Field{
			Type: graphql.Boolean,
		},
		"birthday": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

// UpdateUserGQL __
var UpdateUserGQL = graphql.FieldConfigArgument{
	"username": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"name": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"surname": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"email": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"validated": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Boolean),
	},
	"birthday": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
}
