package main

import "fmt"

func main() {
	fmt.Println(unhex(8))
}
func unhex(C byte) byte {
	switch {
	case '0' <= C && C <= '9':
		return C - '0'
	case 'a' <= C && C <= 'f':
		return C - 'a' + 10
	case 'A' <= C && C <= 'F':
		return C - 'A' + 10
	}
	return 0
}
