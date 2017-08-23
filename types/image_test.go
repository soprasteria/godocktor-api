package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"gopkg.in/mgo.v2/bson"
)

func ExampleImageIdEquals() {
	i := Image{
		ID: bson.ObjectId("1"),
	}
	j := Image{
		ID: bson.ObjectId("1"),
	}
	fmt.Println(i.EqualsInConf(j))
	// Output: true
}

func ExampleImageEquals() {
	i := Image{
		ID:         bson.ObjectId("1"),
		Name:       "temp",
		Variables:  Variables{},
		Ports:      Ports{},
		Volumes:    Volumes{},
		Parameters: Parameters{},
	}
	j := Image{
		ID:         bson.ObjectId("2"),
		Name:       "temp",
		Variables:  Variables{},
		Ports:      Ports{},
		Volumes:    Volumes{},
		Parameters: Parameters{},
	}
	fmt.Println(i.EqualsInConf(j))
	// Output: true
}

func TestImageEquals(t *testing.T) {
	i := Image{
		ID:         bson.ObjectId("1"),
		Name:       "temp",
		Variables:  Variables{},
		Ports:      Ports{},
		Volumes:    Volumes{},
		Parameters: Parameters{},
	}
	j := Image{
		ID:         bson.ObjectId("2"),
		Name:       "temp",
		Variables:  Variables{},
		Ports:      Ports{},
		Volumes:    Volumes{},
		Parameters: Parameters{},
	}
	assert.True(t, i.EqualsInConf(j), "Images should be equals")
}

func TestImageNotEquals(t *testing.T) {
	i := Image{
		ID:   bson.ObjectId("1"),
		Name: "temp",
		Variables: Variables{
			{Name: "var"},
		},
		Ports:      Ports{},
		Volumes:    Volumes{},
		Parameters: Parameters{},
	}
	j := Image{
		ID:   bson.ObjectId("2"),
		Name: "temp",
		Variables: Variables{
			{Name: "notthesame"}, // not the same variables
		},
		Ports:      Ports{},
		Volumes:    Volumes{},
		Parameters: Parameters{},
	}
	assert.False(t, i.EqualsInConf(j), "Images should not be equals")
}

func TestImageAddedPortsEquals(t *testing.T) {
	i := Image{
		ID:        bson.ObjectId("1"),
		Name:      "temp",
		Variables: Variables{},
		Ports: Ports{
			{Internal: 8080, Protocol: "tcp"},
		},
		Volumes:    Volumes{},
		Parameters: Parameters{},
	}
	j := Image{
		ID:        bson.ObjectId("2"),
		Name:      "temp",
		Variables: Variables{},
		Ports: Ports{
			{Internal: 8080, Protocol: "tcp"},
			{Internal: 9090, Protocol: "tcp"},
		},
		Volumes:    Volumes{},
		Parameters: Parameters{},
	}
	assert.False(t, i.EqualsInConf(j), "Images should not be equals")
}

func TestImageRemovedVolumeEquals(t *testing.T) {
	i := Image{
		ID:        bson.ObjectId("1"),
		Name:      "temp",
		Variables: Variables{},
		Ports:     Ports{},
		Volumes: Volumes{
			{Internal: "/tmp", Value: "/external/tmp", Rights: "rw"},
			{Internal: "/data", Value: "/external/data", Rights: "rw"},
		},
		Parameters: Parameters{},
	}
	j := Image{
		ID:        bson.ObjectId("2"),
		Name:      "temp",
		Variables: Variables{},
		Ports:     Ports{},
		Volumes: Volumes{
			{Internal: "/tmp", Value: "/external/tmp", Rights: "rw"},
		},
		Parameters: Parameters{},
	}
	assert.False(t, i.EqualsInConf(j), "Images should not be equals")
}

func TestIsCompatibleWithContainer_okSimple(t *testing.T) {
	i := Image{
		ID:   bson.ObjectId("1"),
		Name: "image1",
		Variables: Variables{
			{Name: "var1", Value: "val1"},
			{Name: "var2", Value: "val2"},
		},
		Ports: Ports{
			{Internal: 1000, Protocol: "tcp"},
			{Internal: 2000, Protocol: "tcp"},
		},
		Volumes: Volumes{
			{Internal: "/tmp", Value: "/external/tmp", Rights: "rw"},
			{Internal: "/data", Value: "/external/data", Rights: "rw"},
		},
		Parameters: Parameters{},
	}
	c := Container{
		Name: "container",
		Variables: VariablesContainer{
			{Name: "var1", Value: "val1"},
			{Name: "var2", Value: "val2"},
			{Name: "var3", Value: "val3"},
		},
		Ports: PortsContainer{
			{Internal: 1000, External: 1000, Protocol: "tcp"},
			{Internal: 2000, External: 2000, Protocol: "tcp"},
		},
		Volumes: VolumesContainer{
			{Internal: "/tmp", External: "/external/tmp", Rights: "rw"},
			{Internal: "/data", External: "/external/data", Rights: "rw"},
		},
		Parameters: ParametersContainer{},
	}

	assert.True(t, i.IsCompatibleWithContainer(c), "Image and container should be compatible")
}

func TestIsCompatibleWithContainer_incompatibleBecauseNewVar(t *testing.T) {
	i := Image{
		ID:   bson.ObjectId("1"),
		Name: "image1",
		Variables: Variables{
			{Name: "var1", Value: ""},
			{Name: "var2", Value: ""},
			{Name: "var3", Value: "val3"},
		},
		Ports: Ports{},
		Volumes: Volumes{
			{Internal: "/tmp", Value: "/external/tmp", Rights: "rw"},
			{Internal: "/data", Value: "/external/data", Rights: "rw"},
		},
		Parameters: Parameters{},
	}
	c := Container{
		Name: "container",
		Variables: VariablesContainer{
			{Name: "var1", Value: "val1"},
			{Name: "var2", Value: "val2"},
		},
		Ports: PortsContainer{},
		Volumes: VolumesContainer{
			{Internal: "/tmp", External: "/external/tmp", Rights: "rw"},
			{Internal: "/data", External: "/external/data", Rights: "rw"},
		},
		Parameters: ParametersContainer{},
	}

	assert.False(t, i.IsCompatibleWithContainer(c), "Image and container should not be compatible because a variable has been added")
}

func TestIsCompatibleWithContainer_compatibleEvenWithNewVolume(t *testing.T) {
	i := Image{
		ID:   bson.ObjectId("1"),
		Name: "image1",
		Variables: Variables{
			{Name: "var1", Value: ""},
			{Name: "var2", Value: ""},
		},
		Ports: Ports{},
		Volumes: Volumes{
			{Internal: "/tmp", Value: "/external/tmp", Rights: "rw"},
			{Internal: "/data", Value: "/external/data", Rights: "rw"},
			{Internal: "/opt", Value: "", Rights: "rw"},
		},
		Parameters: Parameters{},
	}
	c := Container{
		Name: "container",
		Variables: VariablesContainer{
			{Name: "var1", Value: "val1"},
			{Name: "var2", Value: "val2"},
		},
		Ports: PortsContainer{},
		Volumes: VolumesContainer{
			{Internal: "/tmp", External: "/external/tmp", Rights: "rw"},
			{Internal: "/data", External: "/external/data", Rights: "rw"},
		},
		Parameters: ParametersContainer{},
	}

	assert.True(t, i.IsCompatibleWithContainer(c), "Image and container should be compatible even when volume has been added")
}

func TestIsCompatibleWithContainer_incompatibleBecauseRemovedVolume(t *testing.T) {
	i := Image{
		ID:   bson.ObjectId("1"),
		Name: "image1",
		Variables: Variables{
			{Name: "var1", Value: ""},
			{Name: "var2", Value: ""},
		},
		Ports: Ports{},
		Volumes: Volumes{
			{Internal: "/tmp", Value: "/external/tmp", Rights: "rw"},
		},
		Parameters: Parameters{},
	}
	c := Container{
		Name: "container",
		Variables: VariablesContainer{
			{Name: "var1", Value: "val1"},
			{Name: "var2", Value: "val2"},
		},
		Ports: PortsContainer{},
		Volumes: VolumesContainer{
			{Internal: "/tmp", External: "/external/tmp", Rights: "rw"},
			{Internal: "/opt", External: "/external/opt", Rights: "rw"},
		},
		Parameters: ParametersContainer{},
	}

	assert.False(t, i.IsCompatibleWithContainer(c), "Image and container should not be compatible because a volume has been removed")
}

func TestIsCompatibleWithContainer_incompatibleBecauseNewAndRemovedVolume(t *testing.T) {
	i := Image{
		ID:   bson.ObjectId("1"),
		Name: "image1",
		Variables: Variables{
			{Name: "var1", Value: ""},
			{Name: "var2", Value: ""},
		},
		Ports: Ports{},
		Volumes: Volumes{
			{Internal: "/tmp", Value: "/external/tmp", Rights: "rw"},
			{Internal: "/opt", Value: "", Rights: "rw"},
		},
		Parameters: Parameters{},
	}
	c := Container{
		Name: "container",
		Variables: VariablesContainer{
			{Name: "var1", Value: "val1"},
			{Name: "var2", Value: "val2"},
		},
		Ports: PortsContainer{},
		Volumes: VolumesContainer{
			{Internal: "/tmp", External: "/external/tmp", Rights: "rw"},
			{Internal: "/data", External: "/external/data", Rights: "rw"},
		},
		Parameters: ParametersContainer{},
	}

	assert.False(t, i.IsCompatibleWithContainer(c), "Image and container should not be compatible because a volume has been added and another one removed")
}

func TestIsCompatibleWithContainer_compatibleEvenWithDifferentPorts(t *testing.T) {
	i := Image{
		ID:   bson.ObjectId("1"),
		Name: "image1",
		Variables: Variables{
			{Name: "var1", Value: ""},
			{Name: "var2", Value: ""},
		},
		Ports: Ports{
			{Internal: 1000, Protocol: "tcp"},
			{Internal: 2000, Protocol: "tcp"},
			{Internal: 3000, Protocol: "tcp"},
		},
		Volumes: Volumes{
			{Internal: "/tmp", Value: "", Rights: "rw"},
			{Internal: "/data", Value: "", Rights: "rw"},
		},
		Parameters: Parameters{},
	}
	c := Container{
		Name: "container",
		Variables: VariablesContainer{
			{Name: "var1", Value: "val1"},
			{Name: "var2", Value: "val2"},
		},
		Ports: PortsContainer{
			{Internal: 1000, External: 1000, Protocol: "tcp"},
			{Internal: 4000, External: 4000, Protocol: "tcp"},
		},
		Volumes: VolumesContainer{
			{Internal: "/tmp", External: "/external/tmp", Rights: "rw"},
			{Internal: "/data", External: "/external/data", Rights: "rw"},
		},
		Parameters: ParametersContainer{},
	}

	assert.True(t, i.IsCompatibleWithContainer(c), "Image and container should be compatible even when there are different ports")
}

func TestIsCompatibleWithContainer_compatibleEvenWithNoDataAtAll(t *testing.T) {
	i := Image{
		ID:   bson.ObjectId("1"),
		Name: "image1",
	}
	c := Container{
		Name: "container",
	}

	assert.True(t, i.IsCompatibleWithContainer(c), "Image and container does not contain any so should be compatible")
}
