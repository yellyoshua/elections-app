package repository

import (
	"context"
	"errors"
	"time"

	"github.com/yellyoshua/elections-app/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository __
type Repository interface {
	Collection() *mongo.Collection
	Database() *mongo.Database
	UpdateOne(filter interface{}, update map[string]interface{}) error
	FindOne(filter interface{}, dest interface{}) error
	Find(filter interface{}, dest interface{}) error
	FindByID(id primitive.ObjectID, dest interface{}) error
	InsertOne(data interface{}) (primitive.ObjectID, error)
	InsertMany(data []interface{}) error
	Drop() error
}

// Repo _
type Repo struct {
	collection string
}

var db *mongo.Database

// Initialize start database connection
func Initialize(indexes bool) {
	var setup Steps

	setup, db = connect()

	if indexes {
		setup.SetupIndexes()
	}

	if db == nil {
		logger.Fatal("Database connection is nil")
	}
}

// NewRepository __
func NewRepository(collection string) Repository {
	if db == nil {
		logger.Fatal("Need initialize repository connection to database")
	}

	repo := new(Repo)
	repo.collection = collection
	return repo
}

// Database return the database connection
func (r *Repo) Database() *mongo.Database {
	return db
}

// Collection return the database collection session
func (r *Repo) Collection() *mongo.Collection {
	return db.Collection(r.collection)
}

// Drop _
func (r *Repo) Drop() error {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	err := db.Collection(r.collection).Drop(ctx)
	return err
}

// FindByID _
func (r *Repo) FindByID(id primitive.ObjectID, dest interface{}) error {
	var err error

	if err != nil {
		return err
	}

	err = db.Collection(r.collection).FindOne(context.TODO(), bson.M{"_id": id}).Decode(dest)

	if err != nil {
		return err
	}

	return nil
}

// FindOne _
func (r *Repo) FindOne(filter interface{}, dest interface{}) error {
	err := db.Collection(r.collection).FindOne(context.TODO(), filter).Decode(dest)
	if err != nil {
		return err
	}

	return nil
}

// Find __
func (r *Repo) Find(filter interface{}, dest interface{}) error {
	var cursor *mongo.Cursor
	var err error

	ctx := context.Background()
	cursor, err = db.Collection(r.collection).Find(context.TODO(), filter)
	defer cursor.Close(ctx)

	if err != nil {
		return err
	}

	err = cursor.All(ctx, dest)

	if err != nil {
		return err
	}

	return nil
}

// InsertOne _
func (r *Repo) InsertOne(data interface{}) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	result, err := db.Collection(r.collection).InsertOne(ctx, data)

	if err != nil {
		return primitive.ObjectID{}, err
	}

	id, _ := result.InsertedID.(primitive.ObjectID)
	return id, err
}

// InsertMany _
func (r *Repo) InsertMany(data []interface{}) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	_, err := db.Collection(r.collection).InsertMany(ctx, data)
	return err
}

// UpdateOne _
func (r *Repo) UpdateOne(filter interface{}, update map[string]interface{}) error {
	updater, err := db.Collection(r.collection).UpdateOne(context.TODO(), filter, bson.M{"$set": primitive.M(update)})

	if err != nil {
		return err
	}

	if updater.MatchedCount == 0 {
		return errors.New("No matched documents")
	}
	return err
}

// Contraseña pc 1457
