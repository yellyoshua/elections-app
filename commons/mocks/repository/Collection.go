// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package repository

import (
	mock "github.com/stretchr/testify/mock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// Collection is an autogenerated mock type for the Collection type
type Collection struct {
	mock.Mock
}

// Drop provides a mock function with given fields:
func (_m *Collection) Drop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: filter, dest
func (_m *Collection) Find(filter interface{}, dest interface{}) error {
	ret := _m.Called(filter, dest)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, interface{}) error); ok {
		r0 = rf(filter, dest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByID provides a mock function with given fields: id, dest
func (_m *Collection) FindByID(id primitive.ObjectID, dest interface{}) error {
	ret := _m.Called(id, dest)

	var r0 error
	if rf, ok := ret.Get(0).(func(primitive.ObjectID, interface{}) error); ok {
		r0 = rf(id, dest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindOne provides a mock function with given fields: filter, dest
func (_m *Collection) FindOne(filter interface{}, dest interface{}) error {
	ret := _m.Called(filter, dest)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, interface{}) error); ok {
		r0 = rf(filter, dest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertMany provides a mock function with given fields: data
func (_m *Collection) InsertMany(data []interface{}) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func([]interface{}) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertOne provides a mock function with given fields: data
func (_m *Collection) InsertOne(data interface{}) (primitive.ObjectID, error) {
	ret := _m.Called(data)

	var r0 primitive.ObjectID
	if rf, ok := ret.Get(0).(func(interface{}) primitive.ObjectID); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(primitive.ObjectID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateOne provides a mock function with given fields: filter, update
func (_m *Collection) UpdateOne(filter interface{}, update map[string]interface{}) error {
	ret := _m.Called(filter, update)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, map[string]interface{}) error); ok {
		r0 = rf(filter, update)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
