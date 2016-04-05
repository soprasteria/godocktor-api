package sites

import (
	"github.com/soprasteria/godocktor-api/types"
	"github.com/stretchr/testify/mock"
	"gopkg.in/mgo.v2/bson"
)

// MockDocktorSites mocks Docktor sites API
type MockDocktorSites struct {
	mock.Mock
}

// NewMockDocktorSites gets the mock
func NewMockDocktorSites() *MockDocktorSites {
	return &MockDocktorSites{}
}

// Save site into database
func (r *MockDocktorSites) Save(site types.Site) (types.Site, error) {
	args := r.Mock.Called(site)
	return args.Get(0).(types.Site), args.Error(1)
}

// Delete a site in database
func (r *MockDocktorSites) Delete(id bson.ObjectId) (bson.ObjectId, error) {
	args := r.Mock.Called(id)
	return args.Get(0).(bson.ObjectId), args.Error(1)
}

// FindByID get the site by its id
func (r *MockDocktorSites) FindByID(id string) (types.Site, error) {
	args := r.Mock.Called(id)
	return args.Get(0).(types.Site), args.Error(1)
}

// FindByIDBson get the site by its id
func (r *MockDocktorSites) FindByIDBson(id bson.ObjectId) (types.Site, error) {
	args := r.Mock.Called(id)
	return args.Get(0).(types.Site), args.Error(1)
}

// FindAll get all
func (r *MockDocktorSites) FindAll() ([]types.Site, error) {
	args := r.Mock.Called()
	return args.Get(0).([]types.Site), args.Error(1)
}
