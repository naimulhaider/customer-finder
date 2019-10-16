package models

import (
	"reflect"
	"testing"
)

func TestCustomers_SortByUserID(t *testing.T) {
	data := []struct {
		Input          Customers
		ExpectedOutput Customers
	}{
		{
			Input: Customers{
				Customer{
					UserID: 10,
				},
				Customer{
					UserID: 15,
				},
				Customer{
					UserID: 2,
				},
			},
			ExpectedOutput: Customers{
				Customer{
					UserID: 2,
				},
				Customer{
					UserID: 10,
				},
				Customer{
					UserID: 15,
				},
			},
		},
	}

	for _, d := range data {
		d.Input.SortByUserID()
		if !reflect.DeepEqual(d.Input, d.ExpectedOutput) {
			t.Errorf("Expected %v, Got: %v\n", d.ExpectedOutput, d.Input)
		}
	}
}

func TestCustomer_UnmarshalJSON(t *testing.T) {
	data := []struct {
		Input          []byte
		ExpectedOutput Customer
	}{
		{
			Input: []byte(`{"latitude": "52.986375", "user_id": 12, "name": "Christina McArdle", "longitude": "-6.043701"}`),
			ExpectedOutput: Customer{
				UserID: 12,
				Name:   "Christina McArdle",
				Location: Location{
					Latitude:  52.986375,
					Longitude: -6.043701,
				},
			},
		},
		{
			Input: []byte(`{"latitude": "51.92893", "user_id": 1, "name": "Alice Cahill", "longitude": "-10.27699"}`),
			ExpectedOutput: Customer{
				UserID: 1,
				Name:   "Alice Cahill",
				Location: Location{
					Latitude:  51.92893,
					Longitude: -10.27699,
				},
			},
		},
	}

	for _, d := range data {
		customer := Customer{}
		err := customer.UnmarshalJSON(d.Input)
		if err != nil {
			t.Errorf("Unexpected Input: %s: %s\n", d.Input, err.Error())
			continue
		}
		if !reflect.DeepEqual(customer, d.ExpectedOutput) {
			t.Errorf("Expected %v, Got: %v\n", d.ExpectedOutput, d.Input)
		}
	}
}
