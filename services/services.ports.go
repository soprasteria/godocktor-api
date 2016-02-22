package services

import "gopkg.in/mgo.v2/bson"

// Port is a binding between a internal port and an external port
type Port struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Internal    int           `bson:"internal"`
	Protocol    string        `bson:"protocol"` // tcp/udp
	Description string        `bson:"description"`
}

// Ports is a slice of ports
type Ports []Port

// Equals checks that two ports are equals based on some properties
func (a Port) Equals(b Port) bool {
	return a.Internal == b.Internal && a.Protocol == b.Protocol
}

// Equals check that two slices of ports have the same content
func (a Ports) Equals(b Ports) bool {

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
