package database

import "go.mongodb.org/mongo-driver/bson/primitive"

// Profile user
type Profile struct{}

// Permission user
type Permission struct{}

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
	Created   int64              `bson:"created" json:"created"`
	Profile   primitive.ObjectID `bson:"profile" json:"profile"`
}

// Session collection, saved token with userId
type Session struct {
	ID    primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	User  primitive.ObjectID `bson:"user" json:"user,omitempty"`
	Token string             `bson:"token" json:"token"`
}
