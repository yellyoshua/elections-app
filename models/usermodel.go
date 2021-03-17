package models

import gql "github.com/graphql-go/graphql"
import primitive "go.mongodb.org/mongo-driver/bson/primitive"

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

// GraphqlUser _
var GraphqlUser = gql.NewObject(gql.ObjectConfig{
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

// GraphqlCreateUser params graphql create user query
var GraphqlCreateUser = gql.FieldConfigArgument{
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

// GraphqlFindUserByID params graphql
var GraphqlFindUserByID = gql.FieldConfigArgument{
	"id": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
}

// GraphqlFindUserByUsername params graphql
var GraphqlFindUserByUsername = gql.FieldConfigArgument{
	"username": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
}

// GraphqlUpdateUser params graphql update user query
var GraphqlUpdateUser = gql.FieldConfigArgument{
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
