package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	fmt.Println(rangeLoop())
}

func rangeLoop() int {
	marks := []int{
		10, 5, 8, 7, 0,
	}
	sum := 0
	for _, value := range marks {
		sum += value
	}
	fmt.Println(marks[0:2])
	return sum
}
