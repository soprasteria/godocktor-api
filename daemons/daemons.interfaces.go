package daemons

import "gopkg.in/mgo.v2/bson"

// RepoDaemons is the repo for daemons
type RepoDaemons interface {
	// FindByID get the daemon by its id
	FindByID(id string) (Daemon, error)
	// FindByIDBson get the daemon by its (bson representation)
	FindByIDBson(id bson.ObjectId) (Daemon, error)
	// Find get the first daemon with a given name
	Find(name string) (Daemon, error)
	// Find get the first daemon with a given name
	FindAll() ([]Daemon, error)
}
