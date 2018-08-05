// This program prints each line that appears more than
//once preceded by its count.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	i := 1
	for input.Scan() {
		counts[input.Text()]++
		i++
		if i == 6 {
			break
		}
	}
	// NOTE : ignoring potential erros from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
