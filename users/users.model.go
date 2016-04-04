package users

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// User model
type User struct {
	ID          bson.ObjectId   `bson:"_id,omitempty"`
	FirstName   string          `bson:"firstName"`
	LastName    string          `bson:"lastName"`
	DisplayName string          `bson:"displayName"`
	Username    string          `bson:"username"`
	Email       string          `bson:"email"`
	Password    string          `bson:"password"`
	Salt        string          `bson:"salt"`
	Provider    string          `bson:"provider"`
	Role        string          `bson:"role"`
	Created     time.Time       `bson:"created"`
	Updated     time.Time       `bson:"updated"`
	AllowGrant  bool            `bson:"allowGrant"`
	Groups      []bson.ObjectId `bson:"groups"`
	Favorites   []bson.ObjectId `bson:"favorites"`
}
