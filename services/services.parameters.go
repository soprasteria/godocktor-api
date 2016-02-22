package services

import "gopkg.in/mgo.v2/bson"

// Parameter for images ex: CpuShare, etc.
type Parameter struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Name        string        `bson:"name"`
	Value       string        `bson:"value"`
	Description string        `bson:"description"`
}

// Parameters is a slice of parameters
type Parameters []Parameter

// Equals checks that two parameters are equals based on some properties
func (a Parameter) Equals(b Parameter) bool {
	return a.Name == b.Name
}

// Equals check that two slices of parameters have the same content
func (a Parameters) Equals(b Parameters) bool {

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
		if !a[i].Equals(b[i]) {
			return false
		}
	}

	return true
}
