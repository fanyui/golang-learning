package main

import "fmt"

func main() {
	natural := make(chan int)
	sqaures := make(chan int)

	go counter(natural)
	go squarer(sqaures, natural)
	printer(sqaures)
}
func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}
func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
