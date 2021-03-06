package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "Omit trailing new line")
var sep = flag.String("sep", " ", "Seperator")

func main() {
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
