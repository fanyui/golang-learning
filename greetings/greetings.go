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
		return "", errors.New(" Empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
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
