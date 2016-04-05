package types

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Daemon defines a server where services can be deployed
type Daemon struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Created     time.Time     `bson:"created"`
	Name        string        `bson:"name"`
	Protocol    string        `bson:"protocol"`
	Host        string        `bson:"host"`
	Port        int           `bson:"port"`
	Timeout     int           `bson:"timedout"`
	Ca          string        `bson:"ca,omitempty"`
	Cert        string        `bson:"cert,omitempty"`
	Key         string        `bson:"key,omitempty"`
	Volume      string        `bson:"volume,omitempty"`
	Description string        `bson:"description,omitempty"`
	CAdvisorAPI string        `bson:"cadvisorApi,omitempty"`
	User        bson.ObjectId `bson:"user"`
	Site        bson.ObjectId `bson:"site"`
	Variables   Variables     `bson:"variables"`
	Ports       Ports         `bson:"ports"`
	Volumes     Volumes       `bson:"volumes"`
	Parameters  Parameters    `bson:"parameters"`
}

// AddVariable adds a Variable to the Daemon
func (d *Daemon) AddVariable(v Variable) {
	d.Variables = append(d.Variables, v)
}

// AddPort adds a Port to the Daemon
func (d *Daemon) AddPort(p Port) {
	d.Ports = append(d.Ports, p)
}

// AddVolume adds a Volume to the Daemon
func (d *Daemon) AddVolume(v Volume) {
	d.Volumes = append(d.Volumes, v)
}

// AddParameter adds a Parameter to the Daemon
func (d *Daemon) AddParameter(p Parameter) {
	d.Parameters = append(d.Parameters, p)
}
