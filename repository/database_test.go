package repository

import (
	"testing"
)

func TestDatabaseSetup(t *testing.T) {
	// os.Setenv("MONGODB_DATABASE", "golangtest")
	// os.Setenv("MONGODB_URI", "mongodb://root:dbpwd@localhost:27017")
	// var collection string = "demo_indexes"

	// var chanErrs []chan error = []chan error{
	// 	make(chan error),
	// }
	// _, db := connect()

	// cxtTimeout, ctxCancel := context.WithTimeout(context.TODO(), 10*time.Second)
	// defer ctxCancel()

	// col := db.Collection(collection)
	// demoIndexes := []mongo.IndexModel{
	// 	{
	// 		Options: options.Index().SetName("usernameIndex").SetUnique(true).SetDefaultLanguage("en"),
	// 		Keys:    bson.M{"username": 1},
	// 	},
	// 	{
	// 		Options: options.Index().SetName("uuidIndex").SetUnique(true).SetDefaultLanguage("en"),
	// 		Keys:    bson.M{"uuid": -1},
	// 	},
	// }

	// go createIndexes(cxtTimeout, col, demoIndexes, chanErrs[0])

	// for _, c := range chanErrs {
	// 	if err := <-c; err != nil {
	// 		t.Errorf("Failed creating indexes, error: %s", err)
	// 	}
	// }

}
