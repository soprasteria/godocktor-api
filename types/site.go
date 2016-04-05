package types

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Site model
type Site struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Created   time.Time     `bson:"created"`
	Title     string        `bson:"title"`
	Latitude  float64       `bson:"latitude"`
	Longitude float64       `bson:"longitude"`
}
