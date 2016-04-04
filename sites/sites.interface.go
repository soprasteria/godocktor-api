package sites

import "gopkg.in/mgo.v2/bson"

// RepoSites is the repo for sites
type RepoSites interface {
	// Save a site into database
	Save(site Site) (Site, error)
	// Delete a site in database
	Delete(id bson.ObjectId) (bson.ObjectId, error)
	// FindByID get the site by its id
	FindByID(id string) (Site, error)
	// FindByIDBson get the site by its id
	FindByIDBson(id bson.ObjectId) (Site, error)
	// FindAll get all sites
	FindAll() ([]Site, error)
}

