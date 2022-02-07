package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// fmt.Println(Contents("/Users/harisu/Downloads/github/challenge/go/effective_go/loop.go"))
	b()
}
func Contents(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close() // f.Close will run when we're finished.
	var result []byte
	buf := make([]byte, 100)
	for {
		n, err := f.Read(buf[0:])
		result = append(result, buf[0:n]...)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
	}
	return string(result), nil
}

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("Leaving: ", s)
}
func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}
