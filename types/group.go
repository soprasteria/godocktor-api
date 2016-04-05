package types

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

//FileSystems is a slice of FileSystem
type FileSystems []FileSystem

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

// Containers is a slice of Container
type Containers []Container

// AddParameter adds a ParameterContainer to the Container
func (c *Container) AddParameter(p ParameterContainer) {
	c.Parameters = append(c.Parameters, p)
}

// AddPort adds a PortContainer to the Container
func (c *Container) AddPort(p PortContainer) {
	c.Ports = append(c.Ports, p)
}

// AddVariable adds a VariableContainer to the Container
func (c *Container) AddVariable(v VariableContainer) {
	c.Variables = append(c.Variables, v)
}

// AddVolume adds a VolumeContainer to the Container
func (c *Container) AddVolume(v VolumeContainer) {
	c.Volumes = append(c.Volumes, v)
}

// AddJob adds a JobContainer to the Container
func (c *Container) AddJob(j JobContainer) {
	c.Jobs = append(c.Jobs, j)
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
	FileSystems  FileSystems   `bson:"filesystems"`
	Containers   Containers    `bson:"containers"`
	User         bson.ObjectId `bson:"variables"`
}

// AddFileSystem adds a FileSystem to the Group
func (g *Group) AddFileSystem(f FileSystem) {
	g.FileSystems = append(g.FileSystems, f)
}

// AddContainer adds a Container to the Group
func (g *Group) AddContainer(c Container) {
	g.Containers = append(g.Containers, c)
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
