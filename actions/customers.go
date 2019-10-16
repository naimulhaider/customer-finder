package actions

import (
	"github.com/naimulhaider/customer-finder/database"
	"github.com/naimulhaider/customer-finder/models"
)

type Customer struct {
	repo database.Customers
}

func NewCustomer(repo database.Customers) Customer {
	return Customer{
		repo: repo,
	}
}

func (c Customer) GetAllCustomersWithinSearchRadiusFrom(source models.Location, searchRadius float64) (<-chan models.Customer, <-chan error) {

	allCustomers, errors := c.repo.GetAllCustomers()

	customersWithinDistanceFromSource := make(chan models.Customer)

	go func() {
		for customer := range allCustomers {
			if customer.DistanceTo(source) <= searchRadius {
				customersWithinDistanceFromSource <- customer
			}
		}
		close(customersWithinDistanceFromSource)
	}()

	return customersWithinDistanceFromSource, errors
}
