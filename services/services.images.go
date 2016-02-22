package services

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Image defines a docker image
type Image struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	Name       string        `bson:"name"`
	Created    time.Time     `bson:"created"`
	Variables  Variables     `bson:"variables"`
	Ports      Ports         `bson:"ports"`
	Volumes    Volumes       `bson:"volumes"`
	Parameters Parameters    `bson:"parameters"`
	Active     bool
}

// Images is a slice of image
type Images []Image

//EqualsInConf checks that two images are equals in configuration
// It does not check the name for example, but will check ports, variables, parameters and volumes
func (a Image) EqualsInConf(b Image) bool {
	if a.ID == b.ID {
		return true
	}
	return a.Parameters.Equals(b.Parameters) && a.Ports.Equals(b.Ports) && a.Variables.Equals(b.Variables) && a.Volumes.Equals(b.Volumes)
}

// EqualsInConf compare two slices of images by comparing their configuration
func (a Images) EqualsInConf(b Images) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !a[i].EqualsInConf(b[i]) {
			return false
		}
	}

	return true
}

// GetLatestImage gets the last created image for given service
func (s Service) GetLatestImage() (Image, error) {
	var last time.Time
	var image Image

	for _, v := range s.Images {
		created := v.Created
		if v.Created.After(last) {
			last = created
			image = v
		}
	}

	if image.Name == "" {
		return image, errors.New("Did not find any image")
	}

	return image, nil
}

// GetImage returns the image return from the service
func (s Service) GetImage(image string) (Image, error) {
	for _, v := range s.Images {
		if strings.TrimSpace(image) == strings.TrimSpace(v.Name) {
			return v, nil
		}
	}
	return Image{}, fmt.Errorf("Did not find image %v in service %v", image, s.Title)
}

// IsExistingImage checks that image exists in service
func (s Service) IsExistingImage(image string) bool {
	for _, v := range s.Images {
		if strings.TrimSpace(image) == strings.TrimSpace(v.Name) {
			return true
		}
	}

	return false
}
