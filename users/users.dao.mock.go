package users

import (
	"github.com/stretchr/testify/mock"
	"gopkg.in/mgo.v2/bson"
)

// MockDocktorUsers mocks Docktor services API
type MockDocktorUsers struct {
	mock.Mock
}

// NewMockDocktorUsers gets the mock
func NewMockDocktorUsers() *MockDocktorUsers {
	return &MockDocktorUsers{}
}

// Save user into database
func (r *MockDocktorUsers) Save(user User) (User, error) {
	args := r.Mock.Called(user)
	return args.Get(0).(User), args.Error(1)
}

// Delete a user in database
func (r *MockDocktorUsers) Delete(id bson.ObjectId) (bson.ObjectId, error) {
	args := r.Mock.Called(id)
	return args.Get(0).(bson.ObjectId), args.Error(1)
}

// FindByID get the user by its id
func (r *MockDocktorUsers) FindByID(id string) (User, error) {
	args := r.Mock.Called(id)
	return args.Get(0).(User), args.Error(1)
}

// FindByIDBson get the user by its id
func (r *MockDocktorUsers) FindByIDBson(id bson.ObjectId) (User, error) {
	args := r.Mock.Called(id)
	return args.Get(0).(User), args.Error(1)
}

// FindAll get all
func (r *MockDocktorUsers) FindAll() ([]User, error) {
	args := r.Mock.Called()
	return args.Get(0).([]User), args.Error(1)
}

// FindAllByGroupID get all users by a group ID
func (r *MockDocktorUsers) FindAllByGroupID(id bson.ObjectId) ([]User, error) {
	args := r.Mock.Called(id)
	return args.Get(0).([]User), args.Error(1)
}
