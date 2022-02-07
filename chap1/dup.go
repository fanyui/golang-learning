// dup prints the text of each line that appears more
// than once in the standard input preceeded by the count.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	// NOTE: ignoreing the potential errors from input.Error
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t %s \n", n, line)
		}
	}
}
