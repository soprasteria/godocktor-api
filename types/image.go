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
func (i *Image) AddVariable(v *Variable) *Image {
	i.Variables = append(i.Variables, *v)
	return i
}

// AddPort adds a Port to the Image
func (i *Image) AddPort(p *Port) *Image {
	i.Ports = append(i.Ports, *p)
	return i
}

// AddVolume adds a Volume to the Image
func (i *Image) AddVolume(v *Volume) *Image {
	i.Volumes = append(i.Volumes, *v)
	return i
}

// AddParameter adds a Parameter to the Image
func (i *Image) AddParameter(p *Parameter) *Image {
	i.Parameters = append(i.Parameters, *p)
	return i
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

// IsIncludeInConf checks that two images are compatible in configuration
// It does not check the name for example, but will check ports, variables, parameters and volumes
func (i Image) IsIncludeInConf(b Image) bool {
	if i.ID == b.ID {
		return true
	}
	return i.Parameters.IsIncluded(b.Parameters) && i.Ports.IsIncluded(b.Ports) && i.Variables.IsIncluded(b.Variables) && i.Volumes.IsIncluded(b.Volumes)
}
