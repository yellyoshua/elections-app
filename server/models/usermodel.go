package models

import (
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"`

// User user
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name     string             `bson:"name" json:"name"`
	Surname  string             `bson:"surname" json:"surname"`
	Fullname string             `bson:"fullname" json:"fullname"`
	Username string             `bson:"username" json:"username"`
	Email    string             `bson:"email" json:"email"`
	Verified bool               `bson:"verified" json:"verified,omitempty"`
	Password string             `bson:"password" json:"password"`
	Active   bool               `bson:"active" json:"active,omitempty"`
	Created  int64              `bson:"created" json:"created,omitempty"`
}

// UserGQL _
var UserGQL = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"surname": &graphql.Field{
			Type: graphql.String,
		},
		"fullName": &graphql.Field{
			Type: graphql.String,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"verified": &graphql.Field{
			Type: graphql.Boolean,
		},
		"active": &graphql.Field{
			Type: graphql.Boolean,
		},
		"created": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

// CreateUserGQL params graphql create user query
var CreateUserGQL = graphql.FieldConfigArgument{
	"name": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"surname": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"username": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"email": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"password": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}

// FindUserByIDGQL params graphql
var FindUserByIDGQL = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}

// FindUserByUsernameGQL params graphql
var FindUserByUsernameGQL = graphql.FieldConfigArgument{
	"username": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}

// UpdateUserGQL params graphql update user query
var UpdateUserGQL = graphql.FieldConfigArgument{
	"userID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"name": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"surname": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"fullName": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"username": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"email": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"verified": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Boolean),
	},
	"active": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Boolean),
	},
}
