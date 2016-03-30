package groups

import "gopkg.in/mgo.v2/bson"

// RepoGroups is the repo for groups
type RepoGroups interface {
	// Save a group into database
	Save(group Group) (Group, error)
	// Delete a group in database
	Delete(id bson.ObjectId) (bson.ObjectId, error)
	// FindByID get the group by its id
	FindByID(id string) (Group, error)
	// FindByIDBson get the group by its id
	FindByIDBson(id bson.ObjectId) (Group, error)
	// Find get the first group with a given name
	Find(name string) (Group, error)
	// FindAll get all groups
	FindAll() ([]Group, error)
	// FindAllByName get all groups by the give name
	FindAllByName(name string) ([]Group, error)
	// FindAllByRegex get all groups by the regex name
	FindAllByRegex(nameRegex string) ([]Group, error)
	// FindAllWithContainers get all groups that contains a list of containers
	FindAllWithContainers(groupNameRegex string, containersID []string) ([]Group, error)
	// FilterByContainer get all groups matching a regex and a list of containers
	FilterByContainer(groupNameRegex string, service string, containersID []string, imageRegex string) ([]ContainerWithGroup, error)
	// UpdateContainer updates the container from the given group
	UpdateContainer(group Group, container Container) error
}
