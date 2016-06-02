package docktor

import (
	"github.com/soprasteria/godocktor-api/daemons"
	"github.com/soprasteria/godocktor-api/groups"
	"github.com/soprasteria/godocktor-api/services"
	"github.com/soprasteria/godocktor-api/sites"
	"github.com/soprasteria/godocktor-api/users"
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
	Users() users.RepoUsers
	Sites() sites.RepoSites
	Close()
	FindUsedPortsOnDaemon(daemonID string) ([]int, error)
}
