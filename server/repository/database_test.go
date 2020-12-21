package repository

import (
	"context"
	"os"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Test(t *testing.T) {
	os.Setenv("DATABASE_NAME", "golangtest")
	os.Setenv("DATABASE_URI", "mongodb://root:dbpwd@localhost:27017")

	var chanErrs []chan error = []chan error{
		make(chan error),
	}
	db := connect()

	cxtTimeout, ctxCancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer ctxCancel()

	demoCollection := db.Collection("demo")
	demoIndexes := []mongo.IndexModel{
		{
			Options: options.Index().SetName("usernameIndex").SetUnique(true).SetDefaultLanguage("en").SetBackground(true),
			Keys:    bson.M{"username": 1},
		},
	}

	go createIndexes(cxtTimeout, demoCollection, demoIndexes, chanErrs[0])

	for _, c := range chanErrs {
		if err := <-c; err != nil {
			t.Errorf("Failed creating indexes, error: %s", err)
		}
	}

}
