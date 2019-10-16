package main

import (
	"github.com/naimulhaider/customer-finder/actions"
	"github.com/naimulhaider/customer-finder/config"
	"github.com/naimulhaider/customer-finder/database"
	"github.com/naimulhaider/customer-finder/logger"
	"github.com/naimulhaider/customer-finder/models"
)

func main() {
	config.Init()

	customerAction := actions.NewCustomer(database.NewCustomers(config.CustomersFileLocation))

	sourceLocation := models.Location{
		Latitude:  config.SourceLatitude,
		Longitude: config.SourceLongitude,
	}

	customersWithinDistanceFromSource, errors := customerAction.GetAllCustomersWithinSearchRadiusFrom(sourceLocation, config.SearchRadius)

	errorsLogged := logger.LogErrors(errors)
	customersLogged := logger.LogCustomersSortedByID(customersWithinDistanceFromSource)

	<-errorsLogged
	<-customersLogged

	return
}
