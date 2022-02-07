package main

import (
	"fmt"
	"time"
)

func main() {
	list := []int{
		1, 2, 4, 9, 3, 5, 0,
	}
	c := make(chan int)
	go func() {
		fmt.Println("sorting happening here %v", list)
		c <- 1
	}()
	Announce("Public holidays", 10000)
	<-c
	//process the result and sent the feedback to the channel
}
func Announce(message string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		fmt.Println(message)
	}()
}
