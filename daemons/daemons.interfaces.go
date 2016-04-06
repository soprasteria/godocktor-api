package daemons

import (
	"github.com/soprasteria/godocktor-api/types"
	"gopkg.in/mgo.v2/bson"
)

// RepoDaemons is the repo for daemons
type RepoDaemons interface {
	// Save a daemon into database
	Save(daemon types.Daemon) (types.Daemon, error)
	// Delete a daemon in database
	Delete(id bson.ObjectId) (bson.ObjectId, error)
	// FindByID get the daemon by its id
	FindByID(id string) (types.Daemon, error)
	// FindByIDBson get the daemon by its (bson representation)
	FindByIDBson(id bson.ObjectId) (types.Daemon, error)
	// Find get the first daemon with a given name
	Find(name string) (types.Daemon, error)
	// FindAll get all daemons
	FindAll() ([]types.Daemon, error)
	Drop() error
}
