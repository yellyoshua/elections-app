package main

// Profile user
type Profile struct{}

// Permission user
type Permission struct{}

// User user
type User struct {
	ID        string `bson:"_id" json:"_id"`
	Name      string `bson:"name" json:"name"`
	Surname   string `bson:"surname" json:"surname"`
	ProfileID string `bson:"profileId" json:"profileId"`
}

// Session collection, saved token with userId
type Session struct {
	ID     string `bson:"_id" json:"_id"`
	UserID string `bson:"userId" json:"userId"`
	Token  string `bson:"token" json:"token"`
}
