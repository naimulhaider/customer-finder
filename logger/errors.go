package logger

import "fmt"

// LogErrors reads all error data from the channel and then prints it to console
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
