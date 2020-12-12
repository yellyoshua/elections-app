package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Profile user
type Profile struct{}

// Permission user
type Permission struct{}

// Session collection, saved token with userId
type Session struct {
	ID    primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	User  primitive.ObjectID `bson:"user" json:"user,omitempty"`
	Token string             `bson:"token" json:"token"`
}
