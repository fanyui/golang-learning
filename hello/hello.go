package main

import (
	"fmt"

	"go.com/greetings"
)

func main() {
	message := greetings.Hello("Harisu")
	fmt.Println(message)
}
