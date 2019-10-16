package models

import (
	"fmt"
	"math"
	"testing"
)

func TestLocation_DistanceTo(t *testing.T) {
	data := []struct {
		FromLocation Location
		ToLocation   Location
		Distance     float64
	}{
		{
			FromLocation: Location{
				Latitude:  52.986375,
				Longitude: -6.043701,
			},
			ToLocation: Location{
				Latitude:  53.339428,
				Longitude: -6.257664,
			},
			Distance: 41.77,
		},
		{
			FromLocation: Location{
				Latitude:  51.8856167,
				Longitude: -10.4240951,
			},
			ToLocation: Location{
				Latitude:  53.339428,
				Longitude: -6.257664,
			},
			Distance: 324.4,
		},
		{
			FromLocation: Location{
				Latitude:  52.966,
				Longitude: -6.463,
			},
			ToLocation: Location{
				Latitude:  53.339428,
				Longitude: -6.257664,
			},
			Distance: 43.72,
		},
	}

	for _, d := range data {
		output := d.FromLocation.DistanceTo(d.ToLocation)
		if math.Abs(d.Distance-output) > 1 { // check upto 1km precision
			t.Error(fmt.Sprintf("Expected: %f, Got: %f\n", d.Distance, output))
		}
	}
}
