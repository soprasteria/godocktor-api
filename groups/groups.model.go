package groups

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// FileSystem is a filesystem watched by the group
type FileSystem struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Daemon      string        `bson:"daemon"`
	Partition   string        `bson:"partition,omitempty"`
	Description string        `bson:"description"`
}

// ParameterContainer is an env variables given to the creation of the container
type ParameterContainer struct {
	ID    bson.ObjectId `bson:"_id,omitempty"`
	Name  string        `bson:"name"`
	Value string        `bson:"value"`
}

// VariableContainer is a variable for the container
type VariableContainer struct {
	ID    bson.ObjectId `bson:"_id,omitempty"`
	Name  string        `bson:"name"`
	Value string        `bson:"value"`
}

// VolumeContainer is a volume mapped to the container
type VolumeContainer struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Internal string        `bson:"internal"`
	External string        `bson:"external"`
	Rights   string        `bson:"rights"`
}

// JobContainer is a job lunched for the container
type JobContainer struct {
	ID            bson.ObjectId `bson:"_id,omitempty"`
	Name          string        `bson:"name"`
	JobID         string        `bson:"jobId"`
	Description   string        `bson:"description"`
	Result        string        `bson:"result"`
	Status        string        `bson:"status"`
	LastExecution string        `bson:"lastExecution"`
}

// Container is a container associated to the group
type Container struct {
	ID           bson.ObjectId        `bson:"_id,omitempty"`
	Name         string               `bson:"name"`
	Hostname     string               `bson:"hostname"`
	Image        string               `bson:"image"`
	ServiceTitle string               `bson:"serviceTitle"`
	ServiceID    string               `bson:"serviceId"`
	ContainerID  string               `bson:"containerId"`
	Parameters   []ParameterContainer `bson:"parameters"`
	Ports        PortsContainer       `bson:"ports"`
	Variables    []VariableContainer  `bson:"variables"`
	Volumes      []VolumeContainer    `bson:"volumes"`
	Jobs         []JobContainer       `bson:"jobs"`
	DaemonID     string               `bson:"daemonId,omitempty"`
	Active       bool                 `bson:"active"`
}

// Group is a entity (like a project) that gather services instances as containers
type Group struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	Created      time.Time     `bson:"created"`
	Title        string        `bson:"title"`
	Description  string        `bson:"description"`
	PortMinRange int           `bson:"portminrange"`
	PortMaxRange int           `bson:"portmaxrange"`
	Daemon       bson.ObjectId `bson:"daemon"`
	FileSystems  []FileSystem  `bson:"filesystems"`
	Containers   []Container   `bson:"containers"`
	User         bson.ObjectId `bson:"variables"`
}

// ContainerWithGroup is a entity which contains a container, linked to a group
type ContainerWithGroup struct {
	Group     Group
	Container Container
}

// ContainerWithGroupID is an entity which contains a container, linked to a group ID
type ContainerWithGroupID struct {
	Container Container     `bson:"container"`
	ID        bson.ObjectId `bson:"_id,omitempty"`
}
