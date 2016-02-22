package docktor

import (
	"gitlab.cdk.corp.sopra/cdk/godocktor-api/daemons"
	"gitlab.cdk.corp.sopra/cdk/godocktor-api/groups"
	"gitlab.cdk.corp.sopra/cdk/godocktor-api/services"
	"gopkg.in/mgo.v2"
)

// Docktor is the implementation structure to use the API
// It contains API accessing to services, jobs, daemons, etc. + the open session
type Docktor struct {
	services services.RepoServices
	session  Session
	groups   groups.RepoGroups
	daemons  daemons.RepoDaemons
}

type appContext struct {
	db *mgo.Database
}
