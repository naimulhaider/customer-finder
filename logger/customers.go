package logger

import (
	"fmt"

	"github.com/naimulhaider/customer-finder/models"
)

// LogCustomersSortedByID reads all customer data from the channel, sorts the list and then prints it to console
func LogCustomersSortedByID(customers <-chan models.Customer) <-chan bool {
	done := make(chan bool, 1)

	go func() {
		customersForLogging := models.Customers{}

		for customer := range customers {
			customersForLogging = append(customersForLogging, customer)
		}

		customersForLogging.SortByUserID()

		for _, customer := range customersForLogging {
			fmt.Printf("ID: %d, Name: %s\n", customer.UserID, customer.Name)
		}

		done <- true
	}()

	return done
}
