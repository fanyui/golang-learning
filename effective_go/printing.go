package main

import (
	"fmt"
	"os"
)

func main() {
	type T struct {
		a int
		b float64
		c string
		// func (t *T) String() string {
		// 	return fmt.Sprint("%d/%g/%q", t.a, t.b, t.c)
		// }
	}
	t := &T{7, -2.35, "ABC\tDEF"}
	fmt.Printf("Hello %d \n", 23)
	fmt.Fprint(os.Stdout, "hello", 23, "\n")
	fmt.Println("Hello ", 23)
	fmt.Println(fmt.Sprint("hello ", 23))
	fmt.Printf("%v\n", t)
	fmt.Printf("%+v \n", t)
	fmt.Printf("%#v\n", t)
	fmt.Printf("%T\n", t)
	fmt.Printf("path %v", os.Getenv("HOME"))
}
