package daemons

import (
	"github.com/soprasteria/godocktor-api/types"
	"gopkg.in/mgo.v2/bson"
)

// RepoDaemons is the repo for daemons
type RepoDaemons interface {
	// FindByID get the daemon by its id
	FindByID(id string) (types.Daemon, error)
	// FindByIDBson get the daemon by its (bson representation)
	FindByIDBson(id bson.ObjectId) (types.Daemon, error)
	// Find get the first daemon with a given name
	Find(name string) (types.Daemon, error)
	// Find get the first daemon with a given name
	FindAll() ([]types.Daemon, error)
}
