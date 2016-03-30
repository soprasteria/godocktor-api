package daemons

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Repo is the repository for projects
type Repo struct {
	Coll *mgo.Collection
}

// Save a daemon into database
func (r *Repo) Save(daemon Daemon) (Daemon, error) {
	if daemon.ID.Hex() == "" {
		daemon.ID = bson.NewObjectId()
	}

	nb, err := r.Coll.FindId(daemon.ID).Count()
	if err != nil {
		return daemon, err
	}

	if nb != 0 {
		err := r.Coll.UpdateId(daemon.ID, daemon)
		if err != nil {
			return daemon, err
		}
	} else {
		err := r.Coll.Insert(daemon)
		if err != nil {
			return daemon, err
		}
	}
	return daemon, nil
}

// Delete a daemon in database
func (r *Repo) Delete(id bson.ObjectId) (bson.ObjectId, error) {
	err := r.Coll.RemoveId(id)
	if err != nil {
		return id, err
	}
	return id, nil
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

// FindAll get all daemons
func (r *Repo) FindAll() ([]Daemon, error) {
	results := []Daemon{}
	err := r.Coll.Find(bson.M{}).All(&results)
	if err != nil {
		return results, err
	}

	return results, nil
}
