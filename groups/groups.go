package groups

import (
	"github.com/soprasteria/godocktor-api/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Repo is the repository for projects
type Repo struct {
	Coll *mgo.Collection
}

// Drop drops the content of the collection
func (r *Repo) Drop() error {
	return r.Coll.DropCollection()
}

// Save a group into a database
func (r *Repo) Save(group types.Group) (types.Group, error) {
	if group.ID.Hex() == "" {
		group.ID = bson.NewObjectId()
	}

	nb, err := r.Coll.FindId(group.ID).Count()
	if err != nil {
		return group, err
	}

	if nb != 0 {
		err := r.Coll.UpdateId(group.ID, group)
		if err != nil {
			return group, err
		}
	} else {
		err := r.Coll.Insert(group)
		if err != nil {
			return group, err
		}
	}
	return group, nil
}

// Delete a group in database
func (r *Repo) Delete(id bson.ObjectId) (bson.ObjectId, error) {
	err := r.Coll.RemoveId(id)
	if err != nil {
		return id, err
	}
	return id, nil
}

// Find get the first group with a given name
func (r *Repo) Find(name string) (types.Group, error) {
	result := types.Group{}
	err := r.Coll.Find(bson.M{"title": name}).One(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// FindByID get the group by its id
func (r *Repo) FindByID(id string) (types.Group, error) {
	result := types.Group{}
	err := r.Coll.FindId(bson.ObjectIdHex(id)).One(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// FindByIDBson get the group by its id (as a bson object)
func (r *Repo) FindByIDBson(id bson.ObjectId) (types.Group, error) {
	result := types.Group{}
	err := r.Coll.FindId(id).One(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// FindAll get all groups
func (r *Repo) FindAll() ([]types.Group, error) {
	results := []types.Group{}
	err := r.Coll.Find(bson.M{}).All(&results)
	if err != nil {
		return results, err
	}

	return results, nil
}

// FindAllByName get all groups by the give name
func (r *Repo) FindAllByName(name string) ([]types.Group, error) {
	results := []types.Group{}
	err := r.Coll.Find(bson.M{"title": name}).All(&results)
	if err != nil {
		return results, err
	}

	return results, nil
}

// FindAllByRegex get all groups by the regex name
func (r *Repo) FindAllByRegex(nameRegex string) ([]types.Group, error) {
	results := []types.Group{}
	err := r.Coll.Find(bson.M{"title": &bson.RegEx{Pattern: nameRegex}}).All(&results)
	if err != nil {
		return results, err
	}

	return results, nil
}

// FindAllWithContainers get all groups that contains a list of containers
func (r *Repo) FindAllWithContainers(groupNameRegex string, containersID []string) ([]types.Group, error) {
	results := []types.Group{}
	err := r.Coll.Find(
		bson.M{
			"title":                  &bson.RegEx{Pattern: groupNameRegex},
			"containers.containerId": &bson.M{"$in": containersID},
		}).All(&results)

	if err != nil {
		return results, err
	}

	return results, nil
}

// FindAvailablePortsRangeOnDaemon get an available range of <nbConsecutivePorts> for <daemon>
func (r *Repo) FindAvailablePortsRangeOnDaemon(nbConsecutivePorts int, daemon string) (types.PortsRange, error) {

	usedPortsRanges, err := r.FindUsedPortsRangeForDaemon(daemon)
	if err != nil {
		return types.PortsRange{}, err
	}

	return types.FindAvailablePortRange(nbConsecutivePorts, usedPortsRanges)

}

// FindUsedPortsRangeForDaemon get all used ports range for daemon
// Explaination : there is default port range on a group.
func (r *Repo) FindUsedPortsRangeForDaemon(daemon string) ([]types.PortsRange, error) {
	var (
		groups          = []types.Group{}
		usedPortsRanges = []types.PortsRange{}
	)

	err := r.Coll.Find(bson.M{"daemon": daemon}).All(&groups)
	if err != nil {
		return usedPortsRanges, err
	}

	for _, g := range groups {
		if r, err := g.GetPortsRange(); err != nil {
			usedPortsRanges = append(usedPortsRanges, r)
		}
	}

	return usedPortsRanges, nil
}
