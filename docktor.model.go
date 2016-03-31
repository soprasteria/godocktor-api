package docktor

import (
	"github.com/soprasteria/godocktor-api/daemons"
	"github.com/soprasteria/godocktor-api/groups"
	"github.com/soprasteria/godocktor-api/services"
	"github.com/soprasteria/godocktor-api/users"
	"gopkg.in/mgo.v2"
)

// Docktor is the implementation structure to use the API
// It contains API accessing to services, jobs, daemons, etc. + the open session
type Docktor struct {
	services services.RepoServices
	session  Session
	groups   groups.RepoGroups
	daemons  daemons.RepoDaemons
	users    users.RepoUsers
}

type appContext struct {
	db *mgo.Database
}
