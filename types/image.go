package types

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Image defines a docker image
type Image struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	Name       string        `bson:"name"`
	Created    time.Time     `bson:"created"`
	Variables  Variables     `bson:"variables"`
	Ports      Ports         `bson:"ports"`
	Volumes    Volumes       `bson:"volumes"`
	Parameters Parameters    `bson:"parameters"`
	Active     bool
}

// AddVariable adds a Variable to the Image
func (i *Image) AddVariable(v Variable) {
	i.Variables = append(i.Variables, v)
}

// AddPort adds a Port to the Image
func (i *Image) AddPort(p Port) {
	i.Ports = append(i.Ports, p)
}

// AddVolume adds a Volume to the Image
func (i *Image) AddVolume(v Volume) {
	i.Volumes = append(i.Volumes, v)
}

// AddParameter adds a Parameter to the Image
func (i *Image) AddParameter(p Parameter) {
	i.Parameters = append(i.Parameters, p)
}

// Images is a slice of image
type Images []Image

// EqualsInConf checks that two images are equals in configuration
// It does not check the name for example, but will check ports, variables, parameters and volumes
func (i Image) EqualsInConf(b Image) bool {
	if i.ID == b.ID {
		return true
	}
	return i.Parameters.Equals(b.Parameters) && i.Ports.Equals(b.Ports) && i.Variables.Equals(b.Variables) && i.Volumes.Equals(b.Volumes)
}

// EqualsInConf compare two slices of images by comparing their configuration
func (a Images) EqualsInConf(b Images) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !a[i].EqualsInConf(b[i]) {
			return false
		}
	}

	return true
}
