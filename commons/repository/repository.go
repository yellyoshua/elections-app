package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/yellyoshua/elections-app/commons/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// TODO: considerar cambiar el filtro de interface{} -> map[string]interface{}
//  y bson.M{[field]: [filter_value]}

type Collection interface {
	UpdateOne(filter interface{}, update map[string]interface{}) error
	FindOne(filter interface{}, dest interface{}) error
	Find(filter interface{}, dest interface{}) error
	FindByID(id primitive.ObjectID, dest interface{}) error
	InsertOne(data interface{}) (primitive.ObjectID, error)
	InsertMany(data []interface{}) error
	Drop() error
}

type MongoCollection struct {
	current *mongo.Collection
	name    string
	db      *mongo.Database
}

type Client interface {
	Col(collection string) Collection
}

type RepositoryClient struct {
	Conf RepositoryConf
	Db   *mongo.Database
}

type RepositoryConf struct {
	DatabaseURI  string
	DatabaseName string
}

func NoBootstrap(db *mongo.Database) error {
	return nil
}

func New(repository_conf RepositoryConf, database_bootstrap func(*mongo.Database) error) Client {
	db := MongoDatabase(repository_conf.DatabaseURI, repository_conf.DatabaseName)

	if errBootstrap := database_bootstrap(db); errBootstrap != nil {
		logger.Panic("error on bootstrap database repository -> %s", errBootstrap)
	}

	return &RepositoryClient{Db: db, Conf: repository_conf}
}

func NewWithCollection(name string, col *mongo.Collection) Collection {
	return &MongoCollection{
		name:    name,
		db:      &mongo.Database{},
		current: col,
	}
}

func (c *RepositoryClient) Col(collection string) Collection {
	return &MongoCollection{
		name:    collection,
		db:      c.Db,
		current: c.Db.Collection(collection),
	}
}

func (col *MongoCollection) Drop() error {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	return col.current.Drop(ctx)
}

func (col *MongoCollection) FindByID(id primitive.ObjectID, dest interface{}) error {
	filter := bson.M{"_id": id}
	if findError := col.current.FindOne(context.TODO(), filter).Decode(dest); findError != nil {
		return findError
	}

	return nil
}

func (col *MongoCollection) FindOne(filter interface{}, dest interface{}) error {
	if err := col.current.FindOne(context.TODO(), filter).Decode(dest); err != nil {
		return err
	}

	return nil
}

func (col *MongoCollection) Find(filter interface{}, dest interface{}) error {

	finderCtx, cancelFinderCtx := context.WithCancel(context.TODO())
	cursor, findError := col.current.Find(finderCtx, filter)

	defer cancelFinderCtx()

	if findError != nil {
		return findError
	}

	cursorCtx, cancancelCursorCtx := context.WithCancel(context.TODO())
	cursorError := cursor.All(cursorCtx, dest)

	defer cancancelCursorCtx()
	defer cursor.Close(context.TODO())

	if cursorError != nil {
		return cursorError
	}

	return nil
}

func (col *MongoCollection) InsertOne(data interface{}) (primitive.ObjectID, error) {
	ctx, cancel := context.WithCancel(context.TODO())
	result, err := col.current.InsertOne(ctx, data)

	defer cancel()

	if err != nil {
		return primitive.ObjectID{}, err
	}

	id, _ := result.InsertedID.(primitive.ObjectID)
	return id, err
}

func (col *MongoCollection) InsertMany(data []interface{}) error {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	// TODO: verificar si es necesario devolver un array de ID's
	_, err := col.current.InsertMany(ctx, data)
	return err
}

func (col *MongoCollection) UpdateOne(filter interface{}, update map[string]interface{}) error {
	ctx, cancel := context.WithCancel(context.TODO())
	updater, err := col.current.UpdateOne(ctx, filter, bson.M{"$set": primitive.M(update)})

	defer cancel()

	if err != nil {
		return err
	}

	if updater.MatchedCount == 0 {
		return fmt.Errorf("ERROR no matched documents")
	}
	return err
}

/// ---

// package repository

// import (
// 	"context"
// 	"fmt"
// 	"time"

// 	"github.com/yellyoshua/elections-app/commons/logger"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// // Collection __
// type Collection interface {
// 	UpdateOne(filter interface{}, update map[string]interface{}) error
// 	FindOne(filter interface{}, dest interface{}) error
// 	Find(filter interface{}, dest interface{}) error
// 	FindByID(id primitive.ObjectID, dest interface{}) error
// 	InsertOne(data interface{}) (primitive.ObjectID, error)
// 	InsertMany(data []interface{}) error
// 	Drop() error
// }

// type clientStruct struct {
// 	col        *mongo.Collection
// 	collection string
// 	client     *mongo.Database
// }

// // Repository __
// type Repository interface {
// 	Col(collection string) Collection
// 	DatabaseDrop(ctx context.Context) error
// }
// type repositoryStruct struct {
// 	client *mongo.Database
// }

// // Initialize start database connection
// func Initialize(indexes bool) {
// 	setup, db := clientMongodb()

// 	s := &mongo.Collection{}
// 	s.Database()

// 	if indexes {
// 		setup.SetupIndexes()
// 	}

// 	if db == nil {
// 		logger.Fatal("Database connection is nil")
// 	}
// }

// // New _
// func New() Repository {
// 	_, client := clientMongodb()
// 	return NewWithClient(client) // provide real implementation here as argument
// }

// // NewWithClient creates a new Storage client with a custom implementation
// // This is the function you use in your unit tests
// func NewWithClient(client *mongo.Database) Repository {
// 	return &repositoryStruct{
// 		client: client,
// 	}
// }

// func (r *repositoryStruct) Col(collection string) Collection {
// 	// r.client.Collection()
// 	return &clientStruct{
// 		client:     r.client,
// 		col:        r.client.Collection(collection),
// 		collection: collection,
// 	}
// }

// func (r *repositoryStruct) DatabaseDrop(ctx context.Context) error {
// 	return r.client.Drop(ctx)
// }

// func (client *clientStruct) Drop() error {
// 	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
// 	defer cancel()

// 	err := client.col.Drop(ctx)
// 	return err
// }

// func (client *clientStruct) FindByID(id primitive.ObjectID, dest interface{}) error {
// 	var err error

// 	if err != nil {
// 		return err
// 	}

// 	err = client.col.FindOne(context.TODO(), bson.M{"_id": id}).Decode(dest)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (client *clientStruct) FindOne(filter interface{}, dest interface{}) error {
// 	err := client.col.FindOne(context.TODO(), filter).Decode(dest)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (client *clientStruct) Find(filter interface{}, dest interface{}) error {
// 	var cursor *mongo.Cursor
// 	var err error

// 	ctx := context.Background()
// 	cursor, err = client.col.Find(context.TODO(), filter)

// 	if err != nil {
// 		return err
// 	}

// 	err = cursor.All(ctx, dest)
// 	defer cursor.Close(ctx)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (client *clientStruct) InsertOne(data interface{}) (primitive.ObjectID, error) {
// 	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
// 	defer cancel()

// 	result, err := client.col.InsertOne(ctx, data)

// 	if err != nil {
// 		return primitive.ObjectID{}, err
// 	}

// 	id, _ := result.InsertedID.(primitive.ObjectID)
// 	return id, err
// }

// func (client *clientStruct) InsertMany(data []interface{}) error {
// 	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
// 	defer cancel()

// 	_, err := client.col.InsertMany(ctx, data)
// 	return err
// }

// func (client *clientStruct) UpdateOne(filter interface{}, update map[string]interface{}) error {
// 	updater, err := client.col.UpdateOne(context.TODO(), filter, bson.M{"$set": primitive.M(update)})

// 	if err != nil {
// 		return err
// 	}

// 	if updater.MatchedCount == 0 {
// 		return fmt.Errorf("ERROR no matched documents")
// 	}
// 	return err
// }

// func clientMongodb() (Connection, *mongo.Database) {
// 	return connect()
// }

// // Contraseña pc 1457
