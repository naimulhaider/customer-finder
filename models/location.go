package models

import (
	"github.com/naimulhaider/customer-finder/helpers"
)

// Location holds latitude and longitude in degrees
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func NewLocation(lat, lon float64) Location {
	return Location{
		Latitude:  lat,
		Longitude: lon,
	}
}

// DistanceTo returns the distance between the locations in KM
func (l Location) DistanceTo(d Location) float64 {
	return helpers.EarthGeoDistanceBetweenDeg(l.Latitude, l.Longitude, d.Latitude, d.Longitude)
}
