package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("commencing countdown")
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			//process normally
		case <-abort:
			fmt.Println("lauch aborted")
			return
		}
	}
	// launch()
}
