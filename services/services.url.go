package services

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// URL for service
type URL struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	Label   string        `bson:"label"`
	URL     string        `bson:"url"`
	Created time.Time     `bson:"created"`
}
