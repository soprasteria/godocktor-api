package daemons

import (
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

// FindByID get the daemon by its id
func (r *MockDocktorDaemons) FindByID(id string) (Daemon, error) {
	args := r.Mock.Called(id)
	return args.Get(0).(Daemon), args.Error(1)
}

// FindByIDBson get the daemon by its id
func (r *MockDocktorDaemons) FindByIDBson(id bson.ObjectId) (Daemon, error) {
	args := r.Mock.Called(id)
	return args.Get(0).(Daemon), args.Error(1)
}

// Find get the first daemon with a given name
func (r *MockDocktorDaemons) Find(name string) (Daemon, error) {
	args := r.Mock.Called(name)
	return args.Get(0).(Daemon), args.Error(1)
}
