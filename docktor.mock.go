package docktor

import (
	"github.com/soprasteria/godocktor-api/daemons"
	"github.com/soprasteria/godocktor-api/groups"
	"github.com/soprasteria/godocktor-api/services"
	"github.com/soprasteria/godocktor-api/sites"
	"github.com/soprasteria/godocktor-api/users"
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
	users    *users.MockDocktorUsers
	sites    *sites.MockDocktorSites
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
	var mUsers = users.NewMockDocktorUsers()
	var mSites = sites.NewMockDocktorSites()
	var mSession = NewMockSession()
	return &MockDocktor{
		services: mServices,
		groups:   mGroups,
		daemons:  mDaemons,
		users:    mUsers,
		sites:    mSites,
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

// Users mocks the users
func (d *MockDocktor) Users() users.RepoUsers {
	args := d.Mock.Called()
	return args.Get(0).(users.RepoUsers)
}

// Sites mocks the sites
func (d *MockDocktor) Sites() sites.RepoSites {
	args := d.Mock.Called()
	return args.Get(0).(sites.RepoSites)
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

// MockUsers return a mocked users
func (d *MockDocktor) MockUsers() *users.MockDocktorUsers {
	return d.users
}

// MockSites return a mocked sites
func (d *MockDocktor) MockSites() *sites.MockDocktorSites {
	return d.sites
}
