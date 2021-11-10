package greetings

import (
	"errors"
	"fmt"
)

// hello returns a greeting for the named person
func Hello(name string) (string, error) {
	// if no name is provided, return and error with a message
	if name == "" {
		return "", errors.New(" Empty name")
	}
	message := fmt.Sprintf("Hi, %v . Welcome!", name)
	return message, nil
}
