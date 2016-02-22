package services

import "gopkg.in/mgo.v2/bson"

// Volume is a binding between a folder from inside the container to the host machine
type Volume struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Internal    string        `bson:"internal"`         // volume inside the container
	Value       string        `bson:"value"`            // volume outside the contaienr
	Rights      string        `bson:"rights,omitempty"` // ro or rw
	Description string        `bson:"description"`
}

// Volumes is a slice of volumes
type Volumes []Volume

// Equals checks that two volumes are equals based on some properties
func (a Volume) Equals(b Volume) bool {
	return a.Internal == b.Internal && a.Rights == b.Rights
}

// Equals check that two slices of volumes have the same content
func (a Volumes) Equals(b Volumes) bool {

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
