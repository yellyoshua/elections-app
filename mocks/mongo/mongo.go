package mongo

import (
	"context"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"go.mongodb.org/mongo-driver/x/mongo/driver/description"
)

type Collection struct {
	mock.Mock
	client         *mongo.Client
	db             *Database
	name           string
	readConcern    *readconcern.ReadConcern
	writeConcern   *writeconcern.WriteConcern
	readPreference *readpref.ReadPref
	readSelector   description.ServerSelector
	writeSelector  description.ServerSelector
	registry       *bsoncodec.Registry
}

func (c *Collection) Aggregate(ctx context.Context, pipeline interface{}, opts ...*options.AggregateOptions) (*mongo.Cursor, error) {
	ret := c.Called(ctx, pipeline, opts)

	var r0 *mongo.Cursor
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.AggregateOptions) *mongo.Cursor); ok {
		r0 = rf(ctx, pipeline, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.Cursor)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...*options.AggregateOptions) error); ok {
		r1 = rf(ctx, pipeline, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (c *Collection) BulkWrite(ctx context.Context, models []mongo.WriteModel, opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error) {
	ret := c.Called(ctx, models, opts)

	var r0 *mongo.BulkWriteResult
	if rf, ok := ret.Get(0).(func(context.Context, []mongo.WriteModel, ...*options.BulkWriteOptions) *mongo.BulkWriteResult); ok {
		r0 = rf(ctx, models, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.BulkWriteResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []mongo.WriteModel, ...*options.BulkWriteOptions) error); ok {
		r1 = rf(ctx, models, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (c *Collection) Clone(opts ...*options.CollectionOptions) (*mongo.Collection, error) {
	ret := c.Called(opts)

	var r0 *mongo.Collection
	if rf, ok := ret.Get(0).(func(...*options.CollectionOptions) *mongo.Collection); ok {
		r0 = rf(opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.Collection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...*options.CollectionOptions) error); ok {
		r1 = rf(opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (c *Collection) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	ret := c.Called(ctx, filter, opts)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.CountOptions) int64); ok {
		r0 = rf(ctx, filter, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(int64)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...*options.CountOptions) error); ok {
		r1 = rf(ctx, filter, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (c *Collection) Database() *mongo.Database {
	ret := c.Called()

	var r0 *mongo.Database
	if rf, ok := ret.Get(0).(func() *mongo.Database); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.Database)
		}
	}

	return r0
}

type Database struct {
	mock.Mock
}

func (db *Database) Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection {
	ret := db.Called(name, opts)

	var r0 *mongo.Collection
	if rf, ok := ret.Get(0).(func(string, ...*options.CollectionOptions) *mongo.Collection); ok {
		r0 = rf(name, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.Collection)
		}
	}

	return r0
}

// type Database struct {
// 	mock.Mock
// 	client         *Client
// 	name           string
// 	readConcern    *readconcern.ReadConcern
// 	writeConcern   *writeconcern.WriteConcern
// 	readPreference *readpref.ReadPref
// 	readSelector   description.ServerSelector
// 	writeSelector  description.ServerSelector
// 	registry       *Registr
// }

// type Registr struct {
// 	typeEncoders map[reflect.Type]bsoncodec.ValueEncoder
// 	typeDecoders map[reflect.Type]bsoncodec.ValueDecoder

// 	interfaceEncoders []interfaceValueEncoder
// 	interfaceDecoders []interfaceValueDecoder

// 	kindEncoders map[reflect.Kind]bsoncodec.ValueEncoder
// 	kindDecoders map[reflect.Kind]bsoncodec.ValueDecoder

// 	typeMap map[bsontype.Type]reflect.Type

// 	mu sync.RWMutex
// }

// type interfaceValueEncoder struct {
// 	i  reflect.Type
// 	ve bsoncodec.ValueEncoder
// }

// type interfaceValueDecoder struct {
// 	i  reflect.Type
// 	vd bsoncodec.ValueDecoder
// }
