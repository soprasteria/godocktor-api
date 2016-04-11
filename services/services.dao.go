package services

import (
	"fmt"
	"github.com/soprasteria/godocktor-api/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Repo is the repository for services
type Repo struct {
	Coll *mgo.Collection
}

// Save a group into a database
func (r *Repo) Save(service types.Service) (types.Service, error) {
	if service.ID.Hex() == "" {
		service.ID = bson.NewObjectId()
	}

	nb, err := r.Coll.FindId(service.ID).Count()
	if err != nil {
		return service, err
	}

	if nb != 0 {
		err := r.Coll.UpdateId(service.ID, service)
		if err != nil {
			return service, err
		}
	} else {
		err := r.Coll.Insert(service)
		if err != nil {
			return service, err
		}
	}
	return service, nil
}

// Delete a group in database
func (r *Repo) Delete(id bson.ObjectId) (bson.ObjectId, error) {
	err := r.Coll.RemoveId(id)
	if err != nil {
		return id, err
	}
	return id, nil
}

// FindByID the service
func (r *Repo) FindByID(id string) (types.Service, error) {
	result := types.Service{}
	err := r.Coll.FindId(bson.ObjectIdHex(id)).One(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// Find the service by its title, case-insensitive
func (r *Repo) Find(title string) (types.Service, error) {
	result := types.Service{}
	err := r.Coll.Find(bson.M{"title": &bson.RegEx{Pattern: "^" + title + "$", Options: "i"}}).One(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// FindAll get all services by the regex name
func (r *Repo) FindAll() ([]types.Service, error) {
	results := []types.Service{}
	err := r.Coll.Find(bson.M{}).All(&results)
	if err != nil {
		return results, err
	}
	return results, nil
}

// FindAllByRegex get all services by the regex name
func (r *Repo) FindAllByRegex(nameRegex string) ([]types.Service, error) {
	results := []types.Service{}
	err := r.Coll.Find(bson.M{"title": &bson.RegEx{Pattern: nameRegex, Options: "i"}}).All(&results)
	fmt.Println(nameRegex)
	if err != nil {
		return results, err
	}

	return results, nil
}

// IsExist checks that the service exists with given title
func (r *Repo) IsExist(title string) bool {
	_, err := r.Find(title)
	if err != nil {
		return false
	}
	return true
}

func (r *Repo) Drop() error {
	return r.Coll.DropCollection()
}
