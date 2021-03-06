package services

import (
	"testing"
	"time"

	"github.com/soprasteria/godocktor-api/types"
	"github.com/stretchr/testify/assert"
)

func TestGetLastImage(t *testing.T) {
	// Given
	images := []types.Image{
		types.Image{
			Created: time.Date(2005, time.November, 10, 22, 0, 0, 0, time.UTC),
			Name:    "Image2005",
		},
		types.Image{
			Created: time.Date(2016, time.November, 10, 23, 0, 0, 0, time.UTC),
			Name:    "Image2016",
		},
		types.Image{
			Created: time.Date(2015, time.November, 10, 23, 0, 0, 0, time.UTC),
			Name:    "Image2015",
		},
		types.Image{
			Created: time.Date(2010, time.November, 10, 23, 0, 0, 0, time.UTC),
			Name:    "Image2010",
		},
	}
	service := types.Service{
		Images: images,
	}

	// When
	image, err := service.GetLatestImage()

	// Then
	if err != nil {
		t.Log(err.Error())
		t.Error("Can't Get last image")
	}

	assert.Equal(t, "Image2016", image.Name, "Last image should be Image2016")

}
