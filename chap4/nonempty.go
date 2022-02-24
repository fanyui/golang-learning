package main

import (
	"fmt"
)

func main() {
	data := []string{"Hello", "", "fanyui"}
	fmt.Printf("%q", nonempty(data))
	fmt.Printf("%q", data)
}
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}
