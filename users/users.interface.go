package users

import "gopkg.in/mgo.v2/bson"

// RepoUsers is the repo for users
type RepoUsers interface {
	// Save a user into database
	Save(user User) (User, error)
	// Delete a user in database
	Delete(id bson.ObjectId) (bson.ObjectId, error)
	// FindByID get the user by its id
	FindByID(id string) (User, error)
	// FindByIDBson get the user by its id
	FindByIDBson(id bson.ObjectId) (User, error)
	// FindAll get all groups
	FindAll() ([]User, error)
}
