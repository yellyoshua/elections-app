package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/yellyoshua/elections-app/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Client __
type Client interface {
	UpdateOne(filter interface{}, update map[string]interface{}) error
	FindOne(filter interface{}, dest interface{}) error
	Find(filter interface{}, dest interface{}) error
	FindByID(id primitive.ObjectID, dest interface{}) error
	InsertOne(data interface{}) (primitive.ObjectID, error)
	InsertMany(data []interface{}) error
	Drop() error
}

type clientStruct struct {
	col        *mongo.Collection
	collection string
	client     *mongo.Database
}

// Repository __
type Repository interface {
	Col(collection string) Client
	DatabaseDrop(ctx context.Context) error
}
type repositoryStruct struct {
	collection string
	client     *mongo.Database
}

// Initialize start database connection
func Initialize(indexes bool) {
	setup, db := clientMongodb()

	s := &mongo.Collection{}
	s.Database()

	if indexes {
		setup.SetupIndexes()
	}

	if db == nil {
		logger.Fatal("Database connection is nil")
	}
}

// New _
func New() Repository {
	_, client := clientMongodb()
	return NewWithClient(client) // provide real implementation here as argument
}

// NewWithClient creates a new Storage client with a custom implementation
// This is the function you use in your unit tests
func NewWithClient(client *mongo.Database) Repository {
	return &repositoryStruct{
		client: client,
	}
}

func (r *repositoryStruct) Col(collection string) Client {
	// r.client.Collection()
	return &clientStruct{
		client:     r.client,
		col:        r.client.Collection(collection),
		collection: collection,
	}
}

func (r *repositoryStruct) DatabaseDrop(ctx context.Context) error {
	return r.client.Drop(ctx)
}

func (client *clientStruct) Drop() error {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	err := client.col.Drop(ctx)
	return err
}

func (client *clientStruct) FindByID(id primitive.ObjectID, dest interface{}) error {
	var err error

	if err != nil {
		return err
	}

	err = client.col.FindOne(context.TODO(), bson.M{"_id": id}).Decode(dest)

	if err != nil {
		return err
	}

	return nil
}

func (client *clientStruct) FindOne(filter interface{}, dest interface{}) error {
	err := client.col.FindOne(context.TODO(), filter).Decode(dest)
	if err != nil {
		return err
	}

	return nil
}

func (client *clientStruct) Find(filter interface{}, dest interface{}) error {
	var cursor *mongo.Cursor
	var err error

	ctx := context.Background()
	cursor, err = client.col.Find(context.TODO(), filter)
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

func (client *clientStruct) InsertOne(data interface{}) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	result, err := client.col.InsertOne(ctx, data)

	if err != nil {
		return primitive.ObjectID{}, err
	}

	id, _ := result.InsertedID.(primitive.ObjectID)
	return id, err
}

func (client *clientStruct) InsertMany(data []interface{}) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	_, err := client.col.InsertMany(ctx, data)
	return err
}

func (client *clientStruct) UpdateOne(filter interface{}, update map[string]interface{}) error {
	updater, err := client.col.UpdateOne(context.TODO(), filter, bson.M{"$set": primitive.M(update)})

	if err != nil {
		return err
	}

	if updater.MatchedCount == 0 {
		return fmt.Errorf("No matched documents")
	}
	return err
}

func clientMongodb() (Steps, *mongo.Database) {
	return connect()
}

// Contraseña pc 1457
