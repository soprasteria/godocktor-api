package sites

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Repo is the repository for projects
type Repo struct {
	Coll *mgo.Collection
}

// Save a site into a database
func (r *Repo) Save(site Site) (Site, error) {
	if site.ID.Hex() == "" {
		site.ID = bson.NewObjectId()
	}

	nb, err := r.Coll.FindId(site.ID).Count()
	if err != nil {
		return site, err
	}

	if nb != 0 {
		err := r.Coll.UpdateId(site.ID, site)
		if err != nil {
			return site, err
		}
	} else {
		err := r.Coll.Insert(site)
		if err != nil {
			return site, err
		}
	}
	return site, nil
}

// Delete a site in database
func (r *Repo) Delete(id bson.ObjectId) (bson.ObjectId, error) {
	err := r.Coll.RemoveId(id)
	if err != nil {
		return id, err
	}
	return id, nil
}

// FindByID get the site by its id
func (r *Repo) FindByID(id string) (Site, error) {
	result := Site{}
	err := r.Coll.FindId(bson.ObjectIdHex(id)).One(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// FindByIDBson get the site by its id (as a bson object)
func (r *Repo) FindByIDBson(id bson.ObjectId) (Site, error) {
	result := Site{}
	err := r.Coll.FindId(id).One(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// FindAll get all sites
func (r *Repo) FindAll() ([]Site, error) {
	results := []Site{}
	err := r.Coll.Find(bson.M{}).All(&results)
	if err != nil {
		return results, err
	}

	return results, nil
}