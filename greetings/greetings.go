package greetings

import "fmt"

// hello returns a greeting for the named person
func Hello(name string) string {
	// Returns a greeting that embeds the name in message
	message := fmt.Sprintf("Hi, %v . Welcome!", name)
	return message
}
