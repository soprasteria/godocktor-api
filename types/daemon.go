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
	Variables   []Variable    `bson:"variables"`
	Ports       []Port        `bson:"ports"`
	Volumes     []Volume      `bson:"volumes"`
	Parameters  []Parameter   `bson:"parameters"`
}
