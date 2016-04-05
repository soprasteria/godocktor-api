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

// Images is a slice of image
type Images []Image

//EqualsInConf checks that two images are equals in configuration
// It does not check the name for example, but will check ports, variables, parameters and volumes
func (a Image) EqualsInConf(b Image) bool {
	if a.ID == b.ID {
		return true
	}
	return a.Parameters.Equals(b.Parameters) && a.Ports.Equals(b.Ports) && a.Variables.Equals(b.Variables) && a.Volumes.Equals(b.Volumes)
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
