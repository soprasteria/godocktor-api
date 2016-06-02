package types

import (
	"fmt"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

// Volume is a binding between a folder from inside the container to the host machine
type Volume struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Internal    string        `bson:"internal"`         // volume inside the container
	Value       string        `bson:"value"`            // volume outside the contaienr
	Rights      string        `bson:"rights,omitempty"` // ro or rw
	Description string        `bson:"description"`
}

// Print prints a volume as simple string
func (a Volume) Print() string {
	rights := "rw"
	if a.Rights != "" {
		rights = a.Rights
	}
	return a.Value + ":" + a.Internal + ":" + rights
}

func (a Volume) String() string {
	return a.Print()
}

// Check that the volume is well formated
func (a Volume) Check() (bool, error) {
	if a.Rights != "" && a.Rights != "rw" && a.Rights != "ro" {
		return false, fmt.Errorf("Volume rights was %q but should be rw (read-write) or ro (read-only)", a.Rights)
	}

	if a.Internal != "" && !strings.HasPrefix(a.Internal, "/") {
		return false, fmt.Errorf("Volume %q should begin by a '/'. Check that your volume is path like /something/like/this", a.Internal)
	}

	if a.Value != "" && !strings.HasPrefix(a.Value, "/") {
		return false, fmt.Errorf("Volume %q should begin by a '/'. Check that your volume is path like /something/like/this", a.Value)
	}

	return true, nil
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

	if a == nil || b == nil {
		return false
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
