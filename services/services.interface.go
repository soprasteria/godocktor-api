package services

import "gopkg.in/mgo.v2/bson"

//RepoServices is the a repo for services
type RepoServices interface {
	// Save a service
	Save(service Service) (Service, error)
	// Delete a service
	Delete(id bson.ObjectId) (bson.ObjectId, error)
	// FindByID the service
	FindByID(id string) (Service, error)
	// Find the service by its title, case-insensitive
	Find(title string) (Service, error)
	// FindAll get all services
	FindAll() ([]Service, error)
	// FindAllByRegex get all services by the regex name
	FindAllByRegex(nameRegex string) ([]Service, error)
	// IsExist checks that the service exists with given title
	IsExist(title string) bool
}
