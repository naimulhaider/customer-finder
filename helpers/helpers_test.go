package helpers

import (
	"fmt"
	"math"
	"testing"
)

func TestDegToRadian(t *testing.T) {
	data := []struct {
		Deg float64
		Rad float64
	}{
		{
			Deg: 5,
			Rad: 0.0872665,
		},
		{
			Deg: 65,
			Rad: 1.13446,
		},
		{
			Deg: 90,
			Rad: 1.5708,
		},
		{
			Deg: 180,
			Rad: 3.14159,
		},
	}

	for _, d := range data {
		output := DegToRadian(d.Deg)
		if math.Abs(d.Rad-output) > 1e-3 { // check upto 3 digit precision
			t.Error(fmt.Sprintf("Expected: %f, Got: %f\n", d.Rad, output))
		}
	}
}

func TestEarthGeoDistanceBetweenDeg(t *testing.T) {
	data := []struct {
		Lat1     float64
		Lon1     float64
		Lat2     float64
		Lon2     float64
		Distance float64
	}{
		{
			Lat1:     52.986375,
			Lon1:     -6.043701,
			Lat2:     53.339428,
			Lon2:     -6.257664,
			Distance: 41.77,
		},
		{
			Lat1:     51.8856167,
			Lon1:     -10.4240951,
			Lat2:     53.339428,
			Lon2:     -6.257664,
			Distance: 324.4,
		},
		{
			Lat1:     52.966,
			Lon1:     -6.463,
			Lat2:     53.339428,
			Lon2:     -6.257664,
			Distance: 43.72,
		},
	}

	for _, d := range data {
		output := EarthGeoDistanceBetweenDeg(d.Lat1, d.Lon1, d.Lat2, d.Lon2)
		if math.Abs(d.Distance-output) > 1 { // check upto 1km precision
			t.Error(fmt.Sprintf("Expected: %f, Got: %f\n", d.Distance, output))
		}
	}
}
