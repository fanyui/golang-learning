package main

import "fmt"

func main() {
	ages := make(map[string]int)
	ages["alice"] = 21
	ages["harisu"] = 28

	fmt.Printf("%d ==> %d", ages["harisu"], ages["john"])
	delete(ages, "alice")
	ages["fanyui"] += 20
	fmt.Println(ages)
	for name, age := range ages {
		fmt.Printf("%s ==%d \n", name, age)
	}

	bob, ok := ages["bob"]
	if !ok {
		fmt.Printf("Oops bob with age %d is not in the map", bob)
	}
}
