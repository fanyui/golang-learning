package main

import "fmt"

func main() {
	timezones := map[string]int{
		"UTC": 0 * 60 * 60,
		"EST": -5 * 60 * 60,
		"CST": -6 * 60 * 60,
		"MST": -7 * 60 * 60,
		"PST": -8 * 60 * 60,
	}
	fmt.Println(timezones["EST"])
	//checking for presence of a value
	_, present := timezones["EST4"]
	if present {
		fmt.Println("EST is present")
	}
	//deleting an element from the map
	delete(timezones, "PST")
	fmt.Println(timezones)

}
