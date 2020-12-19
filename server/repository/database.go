package repository

import (
	"context"
	"os"
	"time"

	"github.com/yellyoshua/elections-app/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// Connect connection with database
func connect() *mongo.Database {
	var (
		dbName   string = os.Getenv("DATABASE_NAME")
		mongoURI string = os.Getenv("DATABASE_URI")
	)

	// Set client options
	clientOpts := options.Client()

	// Set client URI
	clientURI := clientOpts.ApplyURI(mongoURI)

	logger.Database("Connecting to database")
	defer logger.Database("Connected to database!")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientURI)

	if err != nil {
		logger.DatabaseFatal("Error connection database error: ", err)
		return nil
	}

	// Check the connection
	if err = client.Ping(context.TODO(), nil); err != nil {
		logger.DatabaseFatal("Error ping database error: ", err)
		return nil
	}

	db := client.Database(dbName)

	// setupIndexes(db)

	return db
}

func setupIndexes(db *mongo.Database) {
	var chanErrs []chan error = []chan error{
		make(chan error),
	}

	usersIndexes := []mongo.IndexModel{
		{
			Options: options.Index().SetUnique(true),
			Keys:    bsonx.Doc{{Key: "username", Value: bsonx.String("text")}},
		},
		{
			Options: options.Index().SetUnique(true),
			Keys:    bsonx.Doc{{Key: "username", Value: bsonx.String("text")}},
		},
	}

	defer closeChannels(chanErrs)

	go createIndexes(db.Collection(CollectionUsers), usersIndexes, chanErrs[0])

	for _, c := range chanErrs {
		if err := <-c; err != nil {
			logger.DatabaseFatal("Failed creating indexes, error: %s", err)
		}
	}
	logger.Database("Created indexes!")
}

func createIndexes(col *mongo.Collection, indexes []mongo.IndexModel, c chan error) {
	if existIndexes := checkIndexes(col, indexes); existIndexes {
		c <- nil
		return
	}
	_, err := col.Indexes().CreateMany(context.TODO(), indexes, options.CreateIndexes().SetMaxTime(10*time.Second))
	c <- err
	return
}

func checkIndexes(col *mongo.Collection, indexes []mongo.IndexModel) bool {
	cursor, _ := col.Indexes().List(context.TODO())

	var alreadyExist bool = false

	for cursor.Next(context.TODO()) {
		current, _ := cursor.Current.Values()
		logger.Database("Currents %s", current)

		for _, curr := range current {
			logger.Database("Cursos %s", curr.String())
		}
	}

	return alreadyExist
}

func closeChannels(chans []chan error) {
	for _, c := range chans {
		close(c)
	}
}
