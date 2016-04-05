package types

import (
	"errors"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"time"
)

// Service defines a CDK service
type Service struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Created  time.Time     `bson:"created"`
	Title    string        `bson:"title"`
	Images   Images        `bson:"images"`
	Commands []Command     `bson:"commands"`
	URLs     []URL         `bson:"urls"`
	Jobs     []Job         `bson:"jobs"`
	User     bson.ObjectId `bson:"user"`
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

// GetActiveJobs get active jobs from service.
func (s Service) GetActiveJobs() (jobs []Job) {
	for _, j := range s.Jobs {
		if j.Active {
			jobs = append(jobs, j)
		}
	}
	return
}
