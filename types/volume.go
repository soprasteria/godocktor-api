package types

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

	if len(a) == 0 && len(b) == 0 {
		return true
	}

	if len(a) != len(b) {
		return false
	}

	var aMap = map[string]Volume{}
	for _, v := range a {
		key := v.Internal + ":" + v.Rights
		aMap[key] = v
	}

	for _, v := range b {
		key := v.Internal + ":" + v.Rights
		_, ok := aMap[key]
		if !ok {
			return false
		}
	}

	return true
}

// IsIncluded check that the first slice is included into the second
func (a Volumes) IsIncluded(b Volumes) bool {

	if a == nil && b == nil {
		return true
	}

	if len(a) == 0 && len(b) == 0 {
		return true
	}

	if len(a) > len(b) {
		return false
	}

	var bMap = map[string]Volume{}
	for _, v := range b {
		key := v.Internal + ":" + v.Rights
		bMap[key] = v
	}

	for _, v := range a {
		key := v.Internal + ":" + v.Rights
		_, ok := bMap[key]
		if !ok {
			return false
		}
	}

	return true
}
