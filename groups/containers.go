package groups

import (
	"github.com/soprasteria/godocktor-api/types"
	"gopkg.in/mgo.v2/bson"
)

// FindContainer finds the first container with given containerID
func (r *Repo) FindContainer(groupID bson.ObjectId, containerID string) (types.ContainerWithGroupID, error) {
	result := types.ContainerWithGroupID{}
	// Filter the groups
	filterGroupByTitle := bson.M{"$match": bson.M{"_id": groupID}}
	// Get containers from filtered groups
	getContainers := bson.M{"$unwind": "$containers"}
	// Filter by containers
	filterContainers := bson.M{"$match": bson.M{
		"containers.containerId": &bson.M{"$in": []string{containerID}},
	}}
	// Get ids from containers
	getIds := bson.M{"$project": bson.M{"container": "$containers"}}

	operations := []bson.M{filterGroupByTitle, getContainers, filterContainers, getIds}
	err := r.Coll.Pipe(operations).One(&result)
	return result, err
}

// FindContainerByName finds the first container with given name
// Container Name could have a "/" prefix or not
func (r *Repo) FindContainerByName(groupName string, containerName string) (types.ContainerWithGroupID, error) {
	result := types.ContainerWithGroupID{}

	// match container names like "/ContainerName" or "ContainerName"
	nameRegex := "^\\/?" + containerName + "$"

	operations := []bson.M{
		bson.M{"$match": bson.M{"title": groupName}},
		bson.M{"$unwind": "$containers"},
		bson.M{"$match": &bson.M{"containers.name": &bson.RegEx{Pattern: nameRegex}}},
		bson.M{"$project": bson.M{"container": "$containers"}},
	}
	err := r.Coll.Pipe(operations).One(&result)
	return result, err
}

// FindContainersByService finds all the containers with given service name
func (r *Repo) FindContainersByService(groupName string, serviceName string) ([]types.ContainerWithGroupID, error) {
	result := []types.ContainerWithGroupID{}

	operations := []bson.M{
		bson.M{"$match": bson.M{"title": groupName}},
		bson.M{"$unwind": "$containers"},
		bson.M{"$match": &bson.M{"containers.serviceTitle": serviceName}},
		bson.M{"$project": bson.M{"container": "$containers"}},
	}
	err := r.Coll.Pipe(operations).All(&result)
	return result, err
}

// FilterByContainer get all groups matching a regex and a list of containers
//	db.getCollection('groups').aggregate([
//				{"$match" : {
//						"title": {"$regex" : ".*"}
//						}
//				},
//				{ "$unwind" : "$containers" },
//				{ "$match" : {
//						"containers.containerId": {"$in": ["ID"]},
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

// FilterByContainerAndService returns the data for containers matching a specified group and service
func (r *Repo) FilterByContainerAndService(groupNameRegex string, serviceNameRegex string, containersID []string) (containersWithGroup []types.ContainerWithGroup, err error) {
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
		"containers.serviceTitle": &bson.RegEx{Pattern: serviceNameRegex},
	}}
	// Get ids from containers
	getIds := bson.M{"$project": bson.M{"container": "$containers"}}

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

// FindContainersOnDaemon get all containers that are declared to be run/created on the daemon.
// Can be filtered out with containersID. Will get only containers that are not in this slice.
// db.getCollection('groups').aggregate([
//     {"$match" :
//            {containers: {
//                $elemMatch: {daemonId : "<daemon>" }
//            }}
//     },
//     { "$unwind" : "$containers" },
//     { "$match" : {
// 	        "containers.containerId": {"$nin": ["<containerID1>","<containerID2>"]},
//          "containers.daemonId": "<daemon>"
//     }},
//     { "$project" : {"container" : "$containers"}}
//    ]);
func (r *Repo) FindContainersOnDaemon(daemon types.Daemon, containersID []string) (containersWithGroup []types.ContainerWithGroup, err error) {
	results := []types.ContainerWithGroupID{}

	// Aggregation in 4 steps to get a list of containers
	// Filter the groups
	filterByDaemon := bson.M{"$match": bson.M{
		"containers": &bson.M{"$elemMatch": &bson.M{"daemonId": daemon.ID.Hex()}},
	}}
	// Get containers from filtered groups
	getContainers := bson.M{"$unwind": "$containers"}
	// Containers filtered by daemon and docker containers
	filterContainers := bson.M{"$match": bson.M{
		"containers.containerId": &bson.M{"$nin": containersID},
		"containers.daemonId":    daemon.ID.Hex(),
	}}
	// Get ids from containers
	getIds := bson.M{"$project": bson.M{"container": "$containers"}}

	operations := []bson.M{filterByDaemon, getContainers, filterContainers, getIds}
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

// DeleteContainerByID deletes the container by its docker container ID. The group in which it is, is required
func (r *Repo) DeleteContainerByID(groupID bson.ObjectId, containerID string) error {
	return r.Coll.Update(
		bson.M{"_id": groupID},
		bson.M{"$pull": bson.M{"containerId": containerID}},
	)
}

/*
FindUsedPortsOnDaemon finds ports already used for a daemon
db.getCollection('groups').aggregate(
[
    {'$unwind': '$containers'},
    {'$match': {"containers.daemonId" : "<daemonid>"}},
    {'$unwind' : '$containers.ports'},
    {'$group':
        {
            _id : 0,
            'usedPort' : {'$addToSet': '$containers.ports.external'}
        }
    },
    {'$unwind' : '$usedPort'},
    {'$project' :  {_id: 0, 'usedPort': 1} }
]
)*/
func (r *Repo) FindUsedPortsOnDaemon(daemonID string) ([]int, error) {
	var (
		mgoResults []usedPort
		usedPorts  []int
	)
	// Check if there's already a container with this _id
	operations := []bson.M{
		bson.M{"$unwind": "$containers"},
		bson.M{"$match": bson.M{"containers.daemonId": daemonID}},
		bson.M{"$unwind": "$containers.ports"},
		bson.M{"$group": bson.M{"_id": 0, "port": bson.M{"$addToSet": "$containers.ports.external"}}},
		bson.M{"$unwind": "$port"},
		bson.M{"$project": bson.M{"_id": 0, "port": 1}},
	}
	err := r.Coll.Pipe(operations).All(&mgoResults)
	if err != nil {
		return []int{}, err
	}

	for _, v := range mgoResults {
		usedPorts = append(usedPorts, v.Port)
	}

	return usedPorts, nil
}

type usedPort struct {
	Port int `bson:"port"`
}
