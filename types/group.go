package types

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// FileSystem is a filesystem watched by the group
type FileSystem struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Daemon      bson.ObjectId `bson:"daemon"`
	Partition   string        `bson:"partition,omitempty"`
	Description string        `bson:"description"`
}

//FileSystems is a slice of FileSystem
type FileSystems []FileSystem

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
	User         bson.ObjectId `bson:"user"`
	IsSSO        bool          `bson:"isSSO"`
	Backup       *Backup       `bson:"backup,omitempty"`
}

type Backup struct {
	Created     time.Time `bson:"created"`
	Description string    `bson:"description"`
	Group       Group     `bson:"group"`
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
