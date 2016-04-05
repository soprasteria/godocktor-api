package types

import "gopkg.in/mgo.v2/bson"

// Variable like environment variables (GID of user for example)
type Variable struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Name        string        `bson:"name"`
	Value       string        `bson:"value,omitempty"`
	Description string        `bson:"description"`
}

// Variables is a slice of variables
type Variables []Variable

// Equals checks that two variables are equals based on some properties
func (a Variable) Equals(b Variable) bool {
	return a.Name == b.Name
}

// Equals check that two slices of variables have the same content
func (a Variables) Equals(b Variables) bool {

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
