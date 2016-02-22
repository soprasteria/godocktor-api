package services

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Command for images
type Command struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	Name    string        `bson:"name"`
	Exec    string        `bson:"exec"`
	Role    string        `bson:"role"` // user, admin
	Created time.Time     `bson:"created"`
}
