package database

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/naimulhaider/customer-finder/models"
)

type Customers struct {
	customersFileDir string
}

func NewCustomers(dir string) (c Customers) {
	c.customersFileDir = dir
	return
}

func (c Customers) GetAllCustomers() (<-chan models.Customer, <-chan error) {

	customers := make(chan models.Customer)
	errChan := make(chan error)

	closeChannels := func() {
		close(customers)
		close(errChan)
	}

	customersFile, err := os.Open(c.customersFileDir)
	if err != nil {
		go func() {
			defer closeChannels()
			errChan <- fmt.Errorf("failed to read customers.txt file: %s", err.Error())
		}()
		return customers, errChan
	}

	scanner := bufio.NewScanner(customersFile)

	go func() {
		defer customersFile.Close()
		defer closeChannels()

		var lineNo int
		for scanner.Scan() {
			lineNo++

			customer := models.Customer{}

			err = json.Unmarshal([]byte(scanner.Text()), &customer)
			if err != nil {
				errChan <- fmt.Errorf("malformed JSON in line %d of customers file: %s", lineNo, err.Error())
			} else {
				customers <- customer
			}
		}
	}()

	return customers, errChan
}
