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

	if len(a) == 0 && len(b) == 0 {
		return true
	}

	if len(a) != len(b) {
		return false
	}

	var aMap = map[string]Variable{}
	for _, v := range a {
		aMap[v.Name] = v
	}

	for _, v := range b {
		_, ok := aMap[v.Name]
		if !ok {
			return false
		}
	}

	return true
}

// IsIncluded check that the first slices of variables is included into the second
func (a Variables) IsIncluded(b Variables) bool {

	if a == nil && b == nil {
		return true
	}

	if len(a) == 0 && len(b) == 0 {
		return true
	}

	if len(a) > len(b) {
		return false
	}

	var bMap = map[string]Variable{}
	for _, v := range b {
		bMap[v.Name] = v
	}

	for _, v := range a {
		_, ok := bMap[v.Name]
		if !ok {
			return false
		}
	}
	return true
}
