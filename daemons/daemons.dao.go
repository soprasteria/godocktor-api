package daemons

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Repo is the repository for projects
type Repo struct {
	Coll *mgo.Collection
}

// FindByID get the daemon by its id
func (r *Repo) FindByID(id string) (Daemon, error) {
	result := Daemon{}
	err := r.Coll.FindId(bson.ObjectIdHex(id)).One(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// FindByIDBson get the daemon by its id (as a bson object)
func (r *Repo) FindByIDBson(id bson.ObjectId) (Daemon, error) {
	result := Daemon{}
	err := r.Coll.FindId(id).One(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// Find get the first daemon with a given name (representing the host)
func (r *Repo) Find(name string) (Daemon, error) {
	result := Daemon{}
	err := r.Coll.Find(bson.M{"host": name}).One(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}
