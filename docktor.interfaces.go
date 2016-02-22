package docktor

import (
	"gitlab.cdk.corp.sopra/cdk/godocktor-api/daemons"
	"gitlab.cdk.corp.sopra/cdk/godocktor-api/groups"
	"gitlab.cdk.corp.sopra/cdk/godocktor-api/services"
	"gopkg.in/mgo.v2"
)

// Session is the interface for a docktor sessio
type Session interface {
	SetMode(consistency mgo.Mode, refresh bool)
	Close()
}

//Client is the entrypoint of Docktor API
type Client interface {
	Services() services.RepoServices
	Groups() groups.RepoGroups
	Daemons() daemons.RepoDaemons
	Close()
}
