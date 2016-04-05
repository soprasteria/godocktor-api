package groups

import (
	"github.com/soprasteria/godocktor-api/types"
	"github.com/stretchr/testify/mock"
	"gopkg.in/mgo.v2/bson"
)

// MockDocktorGroups mocks Docktor services API
type MockDocktorGroups struct {
	mock.Mock
}

// NewMockDocktorGroups gets the mock
func NewMockDocktorGroups() *MockDocktorGroups {
	return &MockDocktorGroups{}
}

// Save group into database
func (r *MockDocktorGroups) Save(group types.Group) (types.Group, error) {
	args := r.Mock.Called(group)
	return args.Get(0).(types.Group), args.Error(1)
}

// Delete a group in database
func (r *MockDocktorGroups) Delete(id bson.ObjectId) (bson.ObjectId, error) {
	args := r.Mock.Called(id)
	return args.Get(0).(bson.ObjectId), args.Error(1)
}

// FindByID get the group by its id
func (r *MockDocktorGroups) FindByID(id string) (types.Group, error) {
	args := r.Mock.Called(id)
	return args.Get(0).(types.Group), args.Error(1)
}

// FindByIDBson get the group by its id
func (r *MockDocktorGroups) FindByIDBson(id bson.ObjectId) (types.Group, error) {
	args := r.Mock.Called(id)
	return args.Get(0).(types.Group), args.Error(1)
}

// Find get the first group with a given name
func (r *MockDocktorGroups) Find(name string) (types.Group, error) {
	args := r.Mock.Called(name)
	return args.Get(0).(types.Group), args.Error(1)
}

// FindAll get all
func (r *MockDocktorGroups) FindAll() ([]types.Group, error) {
	args := r.Mock.Called()
	return args.Get(0).([]types.Group), args.Error(1)
}

// FindAllByName get all groups by the give name
func (r *MockDocktorGroups) FindAllByName(name string) ([]types.Group, error) {
	args := r.Mock.Called(name)
	return args.Get(0).([]types.Group), args.Error(1)
}

// FindAllByRegex get all groups by the regex name
func (r *MockDocktorGroups) FindAllByRegex(nameRegex string) ([]types.Group, error) {
	args := r.Mock.Called(nameRegex)
	return args.Get(0).([]types.Group), args.Error(1)
}

// FindAllWithContainers get all groups that contains a list of containers
func (r *MockDocktorGroups) FindAllWithContainers(groupNameRegex string, containersID []string) ([]types.Group, error) {
	args := r.Mock.Called(groupNameRegex, containersID)
	return args.Get(0).([]types.Group), args.Error(1)
}

// FilterByContainer get all groups matching a regex and a list of containers
func (r *MockDocktorGroups) FilterByContainer(groupNameRegex string, service string, containersID []string, imageRegex string) (containersWithGroup []types.ContainerWithGroup, err error) {
	args := r.Mock.Called(groupNameRegex, service, containersID, imageRegex)
	return args.Get(0).([]types.ContainerWithGroup), args.Error(1)
}

// UpdateContainer updates the container from the given group
func (r *MockDocktorGroups) UpdateContainer(group types.Group, container types.Container) error {
	args := r.Mock.Called(group, container)
	return args.Error(0)
}
