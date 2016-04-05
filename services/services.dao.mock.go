package services

import (
	"github.com/soprasteria/godocktor-api/types"
	"github.com/stretchr/testify/mock"
	"gopkg.in/mgo.v2/bson"
)

// MockDocktorServices mocks Docktor services API
type MockDocktorServices struct {
	mock.Mock
}

// NewMockDocktorServices gets the mock
func NewMockDocktorServices() *MockDocktorServices {
	return &MockDocktorServices{}
}

// Save group into database
func (m *MockDocktorServices) Save(service types.Service) (types.Service, error) {
	args := m.Mock.Called(service)
	return args.Get(0).(types.Service), args.Error(1)
}

// Delete a group in database
func (m *MockDocktorServices) Delete(id bson.ObjectId) (bson.ObjectId, error) {
	args := m.Mock.Called(id)
	return args.Get(0).(bson.ObjectId), args.Error(1)
}

// FindByID the service
func (m *MockDocktorServices) FindByID(id string) (types.Service, error) {
	args := m.Mock.Called(id)
	return args.Get(0).(types.Service), args.Error(1)
}

// Find the service by its title, case-insensitive
func (m *MockDocktorServices) Find(title string) (types.Service, error) {
	args := m.Mock.Called(title)
	return args.Get(0).(types.Service), args.Error(1)
}

// FindAll services
func (m *MockDocktorServices) FindAll() ([]types.Service, error) {
	args := m.Mock.Called()
	return args.Get(0).([]types.Service), args.Error(1)
}

// FindAllByRegex the service by regular expression
func (m *MockDocktorServices) FindAllByRegex(title string) ([]types.Service, error) {
	args := m.Mock.Called(title)
	return args.Get(0).([]types.Service), args.Error(1)
}

// IsExist checks that the service exists with given title
func (m *MockDocktorServices) IsExist(title string) bool {
	args := m.Mock.Called(title)
	return args.Bool(0)
}
