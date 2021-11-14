package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// hello returns a greeting for the named person
func Hello(name string) (string, error) {
	// if no name is provided, return and error with a message
	if name == "" {
		return name, errors.New(" Empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	// message := fmt.Sprintf(randomFormat())
	//vecause the name is nor returned in the formated message, it breaks the test for name
	return message, nil
}
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	// a slice of message format
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
func Hellos(names []string) (map[string]string, error) {
	// a map to associate names with messages
	messages := make(map[string]string)
	//loop through the received slice of names, calling
	//the Hello function to get a message for each name
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}
