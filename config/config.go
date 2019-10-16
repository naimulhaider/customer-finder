package config

import "flag"

const (
	CUSTOMERS_FILE_LOCATION = "CUSTOMERS_FILE_LOCATION"
	SOURCE_LAT              = "SOURCE_LAT"
	SOURCE_LON              = "SOURCE_LON"
	SEARCH_RADIUS           = "SEARCH_RADIUS"
)

var (
	CustomersFileLocation string
	SourceLatitude        float64
	SourceLongitude       float64
	SearchRadius          float64 // in KM
)

func Init() {
	ReadFlags()
}

func ReadFlags() {
	flag.StringVar(&CustomersFileLocation, CUSTOMERS_FILE_LOCATION, "data/customers.txt", "The directory of customers.txt file")
	flag.Float64Var(&SourceLatitude, SOURCE_LAT, 53.339428, "The latitude of source")
	flag.Float64Var(&SourceLongitude, SOURCE_LON, -6.257664, "The longitude of source")
	flag.Float64Var(&SearchRadius, SEARCH_RADIUS, 100.0, "The distance (km) from source within which to search for customers")

	flag.Parse()
}
