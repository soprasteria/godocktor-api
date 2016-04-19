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

// Find get the first group with a given name
func (r *Repo) Find(name string) (types.Group, error) {
	result := types.Group{}
	err := r.Coll.Find(bson.M{"title": name}).One(&result)
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

// FilterByContainer get all groups matching a regex and a list of containers
//	db.getCollection('groups').aggregate([
//				{"$match" : {
//						"title": {"$regex" : ".*"}
//						}
//				},
//				{ "$unwind" : "$containers" },
//				{ "$match" : {
//						"containers.containerId": {"$in": ["ID"],
//						"containers.serviceTitle": "Redis",
//						"containers.image" : {"$regex" : "redis:2.*"}
//				}}
//				,
//				{ "$project" : {
//						"_id" : 1,
//						"container" : "$containers",
//						}
//				}
// ])
func (r *Repo) FilterByContainer(groupNameRegex string, service string, containersID []string, imageRegex string) (containersWithGroup []types.ContainerWithGroup, err error) {
	results := []types.ContainerWithGroupID{}

	// Aggregation in 3 steps to get a list of containers id from groups
	// Filter the groups
	filterGroupByTitle := bson.M{"$match": bson.M{
		"title": &bson.RegEx{Pattern: groupNameRegex},
	}}
	// Get containers from filtered groups
	getContainers := bson.M{"$unwind": "$containers"}
	// Filter by containers
	filterContainers := bson.M{"$match": bson.M{
		"containers.containerId":  &bson.M{"$in": containersID},
		"containers.serviceTitle": service,
		"containers.image":        &bson.RegEx{Pattern: imageRegex},
	}}
	// Get ids from containers
	getIds := bson.M{"$project": bson.M{"_id": 1, "container": "$containers"}}

	operations := []bson.M{filterGroupByTitle, getContainers, filterContainers, getIds}
	err = r.Coll.Pipe(operations).All(&results)
	if err != nil {
		return
	}

	// Get group entity for each container
	for _, v := range results {
		group, err := r.FindByIDBson(v.ID)
		if err != nil {
			return []types.ContainerWithGroup{}, err
		}
		crg := types.ContainerWithGroup{
			Group:     group,
			Container: v.Container,
		}
		containersWithGroup = append(containersWithGroup, crg)
	}
	return
}

// update({
//        _id: ObjectId("id"),
//        "containers._id": ObjectId("id")
//    },{
//        $set: {"containers.$": {<container object>}}
//    }
// );
func (r *Repo) updateContainer(group types.Group, container types.Container) error {
	err := r.Coll.Update(
		bson.M{"_id": group.ID, "containers._id": container.ID},
		bson.M{"$set": bson.M{"containers.$": container}},
	)
	return err
}

// SaveContainer saves a container to the given group
func (r *Repo) SaveContainer(group types.Group, container types.Container) error {
	var results []interface{}
	// Check if there's already a container with this _id
	operations := []bson.M{
		bson.M{"$match": bson.M{"_id": group.ID}},
		bson.M{"$unwind": "$containers"},
		bson.M{"$match": bson.M{"containers._id": container.ID}},
		bson.M{"$group": bson.M{"_id": "null", "count": bson.M{"$sum": 1}}},
	}
	err := r.Coll.Pipe(operations).All(&results)
	if err != nil {
		return err
	}
	if len(results) > 0 {
		return r.updateContainer(group, container)
	}

	err = r.Coll.Update(
		bson.M{"_id": group.ID},
		bson.M{"$push": bson.M{"containers": container}},
	)
	return err
}

// Drop drops the content of the collection
func (r *Repo) Drop() error {
	return r.Coll.DropCollection()
}
