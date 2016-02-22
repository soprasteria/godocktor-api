package docktor

import (
	"github.com/soprasteria/godocktor-api/daemons"
	"github.com/soprasteria/godocktor-api/groups"
	"github.com/soprasteria/godocktor-api/services"
	"github.com/stretchr/testify/mock"
	"gopkg.in/mgo.v2"
)

// MockDocktorSession mocks the docktor session
type MockDocktorSession struct {
	mock.Mock
}

//MockDocktor mock Docktor API
type MockDocktor struct {
	services *services.MockDocktorServices
	session  *MockDocktorSession
	groups   *groups.MockDocktorGroups
	daemons  *daemons.MockDocktorDaemons
	mock.Mock
}

// NewMockSession mock the session
func NewMockSession() *MockDocktorSession {
	return &MockDocktorSession{}
}

// NewMockDocktor creates a Docktor API mock
func NewMockDocktor() *MockDocktor {
	var mServices = services.NewMockDocktorServices()
	var mGroups = groups.NewMockDocktorGroups()
	var mDaemons = daemons.NewMockDocktorDaemons()
	var mSession = NewMockSession()
	return &MockDocktor{
		services: mServices,
		groups:   mGroups,
		daemons:  mDaemons,
		session:  mSession,
	}
}

// SetMode sets the mode for the session (mock)
func (m *MockDocktorSession) SetMode(consistency mgo.Mode, refresh bool) {
	m.Mock.Called(consistency, refresh)
}

// Close the session (mock)
func (m *MockDocktorSession) Close() {
	m.Mock.Called()
}

// Close the API (mock)
func (d *MockDocktor) Close() {
	d.Mock.Called()
}

// Services mocks the services
func (d *MockDocktor) Services() services.RepoServices {
	args := d.Mock.Called()
	return args.Get(0).(services.RepoServices)
}

// Groups mocks the groups
func (d *MockDocktor) Groups() groups.RepoGroups {
	args := d.Mock.Called()
	return args.Get(0).(groups.RepoGroups)
}

// Daemons mocks the daemons
func (d *MockDocktor) Daemons() daemons.RepoDaemons {
	args := d.Mock.Called()
	return args.Get(0).(daemons.RepoDaemons)
}

// MockServices return a mocked service
func (d *MockDocktor) MockServices() *services.MockDocktorServices {
	return d.services
}

// MockGroups return a mocked group
func (d *MockDocktor) MockGroups() *groups.MockDocktorGroups {
	return d.groups
}

// MockDaemons return a mocked daemons
func (d *MockDocktor) MockDaemons() *daemons.MockDocktorDaemons {
	return d.daemons
}
