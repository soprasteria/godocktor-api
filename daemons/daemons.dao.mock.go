package daemons

import (
	"github.com/soprasteria/godocktor-api/types"
	"github.com/stretchr/testify/mock"
	"gopkg.in/mgo.v2/bson"
)

// MockDocktorDaemons mocks Daemons services API
type MockDocktorDaemons struct {
	mock.Mock
}

// NewMockDocktorDaemons gets the mock
func NewMockDocktorDaemons() *MockDocktorDaemons {
	return &MockDocktorDaemons{}
}

// Save daemon into database
func (r *MockDocktorDaemons) Save(daemon types.Daemon) (types.Daemon, error) {
	args := r.Mock.Called(daemon)
	return args.Get(0).(types.Daemon), args.Error(1)
}

// Delete a daemon in database
func (r *MockDocktorDaemons) Delete(id bson.ObjectId) (bson.ObjectId, error) {
	args := r.Mock.Called(id)
	return args.Get(0).(bson.ObjectId), args.Error(1)
}

// FindByID get the daemon by its id
func (r *MockDocktorDaemons) FindByID(id string) (types.Daemon, error) {
	args := r.Mock.Called(id)
	return args.Get(0).(types.Daemon), args.Error(1)
}

// FindByIDBson get the daemon by its id
func (r *MockDocktorDaemons) FindByIDBson(id bson.ObjectId) (types.Daemon, error) {
	args := r.Mock.Called(id)
	return args.Get(0).(types.Daemon), args.Error(1)
}

// Find get the first daemon with a given name
func (r *MockDocktorDaemons) Find(name string) (types.Daemon, error) {
	args := r.Mock.Called(name)
	return args.Get(0).(types.Daemon), args.Error(1)
}
