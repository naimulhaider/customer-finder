package logger

import "fmt"

func LogErrors(errors <-chan error) <-chan bool {

	done := make(chan bool, 1)

	go func() {
		for err := range errors {
			fmt.Printf("Encountered error: %s\n", err.Error())
		}
		done <- true
	}()

	return done
}
