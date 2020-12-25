package repository

import (
	"context"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/yellyoshua/elections-app/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Steps _
type Steps interface {
	SetupIndexes()
}

type step struct {
	db *mongo.Database
}

// Connect connection with database
func connect() (Steps, *mongo.Database) {
	var steps = new(step)
	dbName := os.Getenv("DATABASE_NAME")
	mongoURI := os.Getenv("DATABASE_URI")

	cxtTimeout, ctxCancel := context.WithTimeout(context.TODO(), 10*time.Second)

	// Set client options
	clientOpts := options.Client()

	// Set client URI
	clientURI := clientOpts.ApplyURI(mongoURI)

	logger.Database("Connecting to database")
	defer ctxCancel()

	// Connect to MongoDB
	client, err := mongo.Connect(cxtTimeout, clientURI)

	if err != nil {
		logger.DatabaseFatal("Error connection database error: %v", err)
		steps.db = nil
		return steps, nil
	}

	// Check the connection
	if err = client.Ping(cxtTimeout, nil); err != nil {
		logger.DatabaseFatal("Error ping database error: %v", err)
		steps.db = nil
		return steps, nil
	}

	db := client.Database(dbName)

	steps.db = db
	return steps, db
}

// SetupIndexes _
func (s *step) SetupIndexes() {
	ctx, ctxCancel := context.WithTimeout(context.TODO(), 10*time.Second)
	db := s.db

	defer ctxCancel()
	var chanErrs []chan error = []chan error{
		make(chan error),
	}

	usersIndexes := []mongo.IndexModel{
		{
			Options: options.Index().SetName("usernameIndex").SetUnique(true).SetDefaultLanguage("en").SetBackground(true),
			Keys:    bson.M{"username": 1},
		},
	}

	defer closeChannels(ctxCancel, chanErrs)

	go createIndexes(ctx, db.Collection(CollectionUsers), usersIndexes, chanErrs[0])

	for _, c := range chanErrs {
		var err error
		if err = <-c; err != nil {
			logger.DatabaseFatal("Error indexes: %s", err)
		}
	}
	logger.Database("Created/Updated indexes!")
}

func dropAllIndexes(ctx context.Context, col *mongo.Collection) error {
	var err error
	var indexes []interface{}
	var cursor *mongo.Cursor
	cursor, err = col.Indexes().List(ctx)

	if err != nil {
		return errors.Errorf("Listing Indexes %v", err)
	}

	for cursor.Next(ctx) {
		var index interface{}

		err := cursor.Decode(&index)
		if err != nil {
			return errors.Errorf("Indexes cursor %v", err)
		}
		indexes = append(indexes, index)
	}

	if len(indexes) != 0 {
		_, err = col.Indexes().DropAll(ctx)
		if err != nil {
			return errors.Errorf("Dropping Indexes %v", err)
		}
	}

	return err
}

func createIndexes(ctx context.Context, col *mongo.Collection, indexes []mongo.IndexModel, c chan error) {
	if err := dropAllIndexes(ctx, col); err != nil {
		c <- err
		return
	}

	if _, err := col.Indexes().CreateMany(ctx, indexes); err != nil {
		c <- errors.Errorf("Creating Indexes %v", err)
		return
	}

	c <- nil
	return
}

func closeChannels(ctxCancel context.CancelFunc, chans []chan error) {
	for _, c := range chans {
		close(c)
		ctxCancel()
	}
}
