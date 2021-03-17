package repository

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/yellyoshua/elections-app/constants"
	"github.com/yellyoshua/elections-app/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection _
type Connection interface {
	SetupIndexes()
}

type connectionStruct struct {
	db *mongo.Database
}

// Connect connection with database
func connect() (Connection, *mongo.Database) {
	var mongoURI string = os.Getenv("MONGODB_URI")
	var databaseName string

	cxtTimeout, ctxCancel := context.WithTimeout(context.TODO(), 5*time.Second)

	// Set URI to client options
	clientURI := options.Client().ApplyURI(mongoURI)

	logger.Info("Connecting to database")
	defer ctxCancel()

	// Connect to MongoDB
	client, err := mongo.Connect(cxtTimeout, clientURI)

	if err != nil {
		logger.Fatal("Error when trying connect to mongo database -> %v", err)
		return &connectionStruct{db: nil}, nil
	}

	// Checking the connection to database
	if err = client.Ping(cxtTimeout, nil); err != nil {
		logger.Fatal("Error ping database error: %v", err)
		return &connectionStruct{nil}, nil
	}

	if len(os.Getenv("MONGODB_DATABASE")) == 0 {
		databaseName = constants.DefaultDatabase
	} else {
		databaseName = os.Getenv("MONGODB_DATABASE")
	}

	db := client.Database(databaseName)

	return &connectionStruct{db}, db
}

// SetupIndexes _
func (s *connectionStruct) SetupIndexes() {
	ctx, ctxCancel := context.WithTimeout(context.TODO(), 10*time.Second)
	db := s.db

	defer ctxCancel()
	var chanErrs []chan error = []chan error{
		make(chan error),
	}

	usersIndexes := []mongo.IndexModel{
		{
			Options: options.Index().SetName("usernameIndex").SetUnique(true).SetDefaultLanguage("en"),
			Keys:    bson.M{"username": 1},
		},
	}

	defer closeChannels(ctxCancel, chanErrs)

	go createIndexes(ctx, db.Collection(constants.CollectionUsers), usersIndexes, chanErrs[0])

	for _, c := range chanErrs {
		var err error
		if err = <-c; err != nil {
			logger.Fatal("Error indexes: %s", err)
		}
	}
	logger.Info("Database created/updated indexes!")
}

func createIndexes(ctx context.Context, col *mongo.Collection, indexes []mongo.IndexModel, c chan error) {
	if err := dropAllIndexes(ctx, col); err != nil {
		c <- err
		return
	}

	if _, err := col.Indexes().CreateMany(ctx, indexes); err != nil {
		c <- fmt.Errorf("ERROR creating indexes -> %v", err)
		return
	}

	c <- nil
}

func closeChannels(ctxCancel context.CancelFunc, chans []chan error) {
	for _, c := range chans {
		close(c)
		ctxCancel()
	}
}

func dropAllIndexes(ctx context.Context, col *mongo.Collection) error {
	var err error
	var indexes []interface{}
	var cursor *mongo.Cursor
	cursor, err = col.Indexes().List(ctx)

	if err != nil {
		return fmt.Errorf("ERROR listing indexes -> %v", err)
	}

	for cursor.Next(ctx) {
		var index interface{}

		err := cursor.Decode(&index)
		if err != nil {
			return fmt.Errorf("ERROR with indexes cursor -> %v", err)
		}
		indexes = append(indexes, index)
	}

	if len(indexes) != 0 {
		_, err = col.Indexes().DropAll(ctx)
		if err != nil {
			return fmt.Errorf("ERROR dropping indexes -> %v", err)
		}
	}

	return err
}
