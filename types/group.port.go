package types

import "gopkg.in/mgo.v2/bson"

// PortContainer defines a binding betweend an external and an internal port
type PortContainer struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Internal int           `bson:"internal"`
	External int           `bson:"external"`
	Protocol string        `bson:"protocol"`
}

// PortsContainer is a slice of PortContainer
type PortsContainer []PortContainer

// GetExternalPort search the external port bind to a given internalPort
func (ports PortsContainer) GetExternalPort(internalPort int) int {
	for _, p := range ports {
		if p.Internal == internalPort {
			return p.External
		}
	}
	return 0
}
