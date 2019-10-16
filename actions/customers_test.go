package actions

import (
	"fmt"
	"reflect"
	"sync"
	"testing"

	"github.com/naimulhaider/customer-finder/models"

	"github.com/naimulhaider/customer-finder/database"
)

func TestCustomer_GetAllCustomersWithinSearchRadiusFrom(t *testing.T) {

	customerAction := NewCustomer(database.NewCustomers("../data/customers_test_data.txt"))

	customers, errors := customerAction.GetAllCustomersWithinSearchRadiusFrom(models.Location{
		Latitude:  53.339428,
		Longitude: -6.257664,
	}, 100.0)

	w := sync.WaitGroup{}
	w.Add(2)

	customerList := models.Customers{}
	errorList := []error{}

	go func() {
		for customer := range customers {
			customerList = append(customerList, customer)
		}
		w.Done()
	}()

	go func() {
		for err := range errors {
			errorList = append(errorList, err)
		}
		w.Done()
	}()

	w.Wait()

	expectedCustomers := models.Customers{
		models.Customer{
			UserID: 12,
			Name:   "Christina McArdle",
			Location: models.Location{
				Latitude:  52.986375,
				Longitude: -6.043701,
			},
		},
		models.Customer{
			UserID: 13,
			Name:   "Olive Ahearn",
			Location: models.Location{
				Latitude:  53,
				Longitude: -7,
			},
		},
	}

	expectedErrors := []error{
		fmt.Errorf("malformed JSON in line 3 of customers file: invalid character 's' looking for beginning of object key string"),
	}

	if !reflect.DeepEqual(expectedCustomers, customerList) {
		t.Errorf("Customers mismatch, Expected: %v, Got: %v\n", expectedCustomers, customerList)
	}

	if !reflect.DeepEqual(expectedErrors, errorList) {
		t.Errorf("Errors mismatch, Expected: %v, Got: %v\n", expectedErrors, errorList)
	}

}
