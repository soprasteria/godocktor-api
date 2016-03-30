package groups

import (
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
func (r *MockDocktorGroups) Save(group Group) (Group, error) {
	args := r.Mock.Called(group)
	return args.Get(0).(Group), args.Error(1)
}

// Delete a group in database
func (r *MockDocktorGroups) Delete(id bson.ObjectId) (bson.ObjectId, error) {
	args := r.Mock.Called(id)
	return args.Get(0).(bson.ObjectId), args.Error(1)
}

// FindByID get the group by its id
func (r *MockDocktorGroups) FindByID(id string) (Group, error) {
	args := r.Mock.Called(id)
	return args.Get(0).(Group), args.Error(1)
}

// FindByIDBson get the group by its id
func (r *MockDocktorGroups) FindByIDBson(id bson.ObjectId) (Group, error) {
	args := r.Mock.Called(id)
	return args.Get(0).(Group), args.Error(1)
}

// Find get the first group with a given name
func (r *MockDocktorGroups) Find(name string) (Group, error) {
	args := r.Mock.Called(name)
	return args.Get(0).(Group), args.Error(1)
}

// FindAll get all
func (r *MockDocktorGroups) FindAll() ([]Group, error) {
	args := r.Mock.Called()
	return args.Get(0).([]Group), args.Error(1)
}

// FindAllByName get all groups by the give name
func (r *MockDocktorGroups) FindAllByName(name string) ([]Group, error) {
	args := r.Mock.Called(name)
	return args.Get(0).([]Group), args.Error(1)
}

// FindAllByRegex get all groups by the regex name
func (r *MockDocktorGroups) FindAllByRegex(nameRegex string) ([]Group, error) {
	args := r.Mock.Called(nameRegex)
	return args.Get(0).([]Group), args.Error(1)
}

// FindAllWithContainers get all groups that contains a list of containers
func (r *MockDocktorGroups) FindAllWithContainers(groupNameRegex string, containersID []string) ([]Group, error) {
	args := r.Mock.Called(groupNameRegex, containersID)
	return args.Get(0).([]Group), args.Error(1)
}

// FilterByContainer get all groups matching a regex and a list of containers
func (r *MockDocktorGroups) FilterByContainer(groupNameRegex string, service string, containersID []string, imageRegex string) (containersWithGroup []ContainerWithGroup, err error) {
	args := r.Mock.Called(groupNameRegex, service, containersID, imageRegex)
	return args.Get(0).([]ContainerWithGroup), args.Error(1)
}

// UpdateContainer updates the container from the given group
func (r *MockDocktorGroups) UpdateContainer(group Group, container Container) error {
	args := r.Mock.Called(group, container)
	return args.Error(0)
}
