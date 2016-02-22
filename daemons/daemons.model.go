package daemons

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Variable is used to have common variables on all services deployed on this daemon
type Variable struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Name        string        `bson:"name"`
	Value       string        `bson:"value,omitempty"`
	Description string        `bson:"description"`
}

// Parameter is used to have common parameters on all services deployed on this daemon
type Parameter struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Name        string        `bson:"name"`
	Value       string        `bson:"value,omitempty"`
	Description string        `bson:"description"`
}

// Port is used to have common ports on all services deployed on this daemon
type Port struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Internal    int           `bson:"internal"`
	Protocol    string        `bson:"protocol"` // tcp, udp
	Description string        `bson:"description"`
}

// Volume is used to have common volumes on all services deployed on this daemon
type Volume struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Internal    string        `bson:"internal"`
	Value       string        `bson:"value,omitempty"`  // default value, ex : /etc/localtime
	Rights      string        `bson:"rights,omitempty"` // ro or rw
	Description string        `bson:"description"`
}

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
	Variables   []Variable    `bson:"variables"`
	Ports       []Port        `bson:"ports"`
	Volumes     []Volume      `bson:"volumes"`
	Parameters  []Parameter   `bson:"parameters"`
}
