package services

//RepoServices is the a repo for services
type RepoServices interface {
	// FindByID the service
	FindByID(id string) (Service, error)
	// Find the service by its title, case-insensitive
	Find(title string) (Service, error)
	// IsExist checks that the service exists with given title
	IsExist(title string) bool
}
