// prints the command line arguements just the way tey were passed.
package main

import (
	"fmt"
	"os"
)

func main() {
	// var s, sep string
	if len(os.Args) == 1 {
		fmt.Print(" no args")
	}
	// for i := 1; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// 	sep = " "
	// }
	// fmt.Println(s)
	fmt.Print(os.Args[1:])

	// fmt.Print(strings.Join(os.Args[1:], " "))
}
