package services

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Service defines a CDK service
type Service struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Created  time.Time     `bson:"created"`
	Title    string        `bson:"title"`
	Images   Images        `bson:"images"`
	Commands []Command     `bson:"commands"`
	URLs     []URL         `bson:"urls"`
	Jobs     []Job         `bson:"jobs"`
	User     bson.ObjectId `bson:"user"`
}

// Repo is the repository for services
type Repo struct {
	Coll *mgo.Collection
}

// FindByID the service
func (r *Repo) FindByID(id string) (Service, error) {
	result := Service{}
	err := r.Coll.FindId(bson.ObjectIdHex(id)).One(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// Find the service by its title, case-insensitive
func (r *Repo) Find(title string) (Service, error) {
	result := Service{}
	err := r.Coll.Find(bson.M{"title": &bson.RegEx{Pattern: "^" + title + "$", Options: "i"}}).One(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// IsExist checks that the service exists with given title
func (r *Repo) IsExist(title string) bool {
	_, err := r.Find(title)
	if err != nil {
		return false
	}
	return true
}
