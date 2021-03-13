package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Profile user
type Profile struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
}

// Permission user
type Permission struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
}

// Session collection, saved token with userId
type Session struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Token string             `bson:"token" json:"token"`
}
