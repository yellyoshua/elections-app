package repository

import (
	"context"

	"github.com/yellyoshua/elections-app/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository __
type Repository interface {
	FindOne(filter interface{}, dest interface{}) error
	Find(filter interface{}, dest interface{}) error
	InsertOne(data interface{})
}

// Repo _
type Repo struct {
	collection string
}

var db *mongo.Database

// InitializeMock start mock database connection
func InitializeMock() {
	db = connect()

	if db == nil {
		logger.DatabaseFatal("Database connection is nil")
	}
}

// Initialize start database connection
func Initialize() {
	db = connect()

	if db == nil {
		logger.DatabaseFatal("Database connection is nil")
	}
}

// NewRepository __
func NewRepository(collection string) Repository {
	if db == nil {
		logger.DatabaseFatal("Need initialize repository connection to database")
	}

	repo := new(Repo)
	repo.collection = collection
	return repo
}

// FindOne _
func (r *Repo) FindOne(filter interface{}, dest interface{}) error {
	err := db.Collection(r.collection).FindOne(context.TODO(), filter).Decode(&dest)
	if err != nil {
		return err
	}

	return nil
}

// Find __
func (r *Repo) Find(filter interface{}, dest interface{}) error {

	cursor, err := db.Collection(r.collection).Find(context.TODO(), filter)

	if err != nil {
		return err
	}

	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&dest)
		if err != nil {
			return err
		}
	}

	return nil
}

// InsertOne _
func (r *Repo) InsertOne(data interface{}) {
	db.Collection(r.collection).InsertOne(context.TODO(), data)
}

// Contraseña pc 1457
