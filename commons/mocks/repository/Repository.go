// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package repository

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	repository "github.com/yellyoshua/elections-app/commons/repository"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Col provides a mock function with given fields: collection
func (_m *Repository) Col(collection string) repository.Collection {
	ret := _m.Called(collection)

	var r0 repository.Collection
	if rf, ok := ret.Get(0).(func(string) repository.Collection); ok {
		r0 = rf(collection)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(repository.Collection)
		}
	}

	return r0
}

// DatabaseDrop provides a mock function with given fields: ctx
func (_m *Repository) DatabaseDrop(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
