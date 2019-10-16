package models

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
)

type Customers []Customer

func (c Customers) SortByUserID() {
	sort.Slice(c, func(i int, j int) bool {
		return c[i].UserID < c[j].UserID
	})
}

type Customer struct {
	UserID int64  `json:"user_id"`
	Name   string `json:"name"`
	Location
}

// UnmarshalJSON
func (c *Customer) UnmarshalJSON(data []byte) error {
	jsonData := struct {
		UserID    int64  `json:"user_id"`
		Name      string `json:"name"`
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
	}{}

	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		return err
	}

	latitude, err := strconv.ParseFloat(jsonData.Latitude, 64)
	if err != nil {
		return fmt.Errorf("failed to parse latitude as float64: %s", err.Error())
	}

	longitude, err := strconv.ParseFloat(jsonData.Longitude, 64)
	if err != nil {
		return fmt.Errorf("failed to parse latitude as float64: %s", err.Error())
	}

	c.UserID = jsonData.UserID
	c.Name = jsonData.Name
	c.Latitude = latitude
	c.Longitude = longitude

	return nil
}
