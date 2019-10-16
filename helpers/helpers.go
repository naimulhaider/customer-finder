package helpers

import (
	"math"
)

const (
	EARTH_RADIUS = 6371 // km
)

func DegToRadian(deg float64) float64 {
	return deg * math.Pi / 180
}

// EarthGeoDistanceBetweenDeg takes latitudes and longitudes of two points on earth in DEGREES and returns the distance in KM using the haversine formula
func EarthGeoDistanceBetweenDeg(lat1, lon1, lat2, lon2 float64) float64 {
	phi1 := DegToRadian(lat1)

	phi2 := DegToRadian(lat2)

	dPhi := DegToRadian(lat2 - lat1)

	dLambda := DegToRadian(lon2 - lon1)

	a := math.Sin(dPhi/2)*math.Sin(dPhi/2) + math.Cos(phi1)*math.Cos(phi2)*math.Sin(dLambda/2)*math.Sin(dLambda/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return EARTH_RADIUS * c
}
