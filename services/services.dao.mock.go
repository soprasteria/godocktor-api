package services

import "github.com/stretchr/testify/mock"

// MockDocktorServices mocks Docktor services API
type MockDocktorServices struct {
	mock.Mock
}

// NewMockDocktorServices gets the mock
func NewMockDocktorServices() *MockDocktorServices {
	return &MockDocktorServices{}
}

// FindByID the service
func (m *MockDocktorServices) FindByID(id string) (Service, error) {
	args := m.Mock.Called(id)
	return args.Get(0).(Service), args.Error(1)
}

// Find the service by its title, case-insensitive
func (m *MockDocktorServices) Find(title string) (Service, error) {
	args := m.Mock.Called(title)
	return args.Get(0).(Service), args.Error(1)
}

// FindAllByRegex the service by regular expression
func (m *MockDocktorServices) FindAllByRegex(title string) ([]Service, error) {
	args := m.Mock.Called(title)
	return args.Get(0).([]Service), args.Error(1)
}

// IsExist checks that the service exists with given title
func (m *MockDocktorServices) IsExist(title string) bool {
	args := m.Mock.Called(title)
	return args.Bool(0)
}
