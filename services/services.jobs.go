package services

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Job for service
type Job struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Name        string        `bson:"name"`
	Value       string        `bson:"value"`    // ":port" if type = url, "unix command" if type= exec
	Interval    string        `bson:"interval"` // cron format
	Type        string        `bson:"type"`     // url/exec
	Description string        `bson:"description"`
	Active      bool          `bson:"active"`
	Created     time.Time     `bson:"created"`
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

//func (j Job)
