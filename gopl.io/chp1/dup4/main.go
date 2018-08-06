// This program prints each line that appears more than
//once preceded by its count.
// The lines come from standard input or given files in command-line
// It prints also the name of files
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filenames := make(map[string]string)
	files := os.Args[1:]

	if len(files) == 0 {
		countLinesStdin(os.Stdin, counts, filenames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Erreur : %v\n", err)
				continue
			}
			countLinesFiles(f, counts, filenames)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, filenames[line])
		}
	}
}

func countLinesStdin(f *os.File, counts map[string]int, filenames map[string]string) {
	input := bufio.NewScanner(f)
	i := 1
	for input.Scan() {
		counts[input.Text()]++
		if filenames[input.Text()] == "" {
			filenames[input.Text()] = os.Stdin.Name()
		}
		i++
		if i > 5 {
			break
		}
	}
	// NOTE : ignoring potential erros from input.Err()
}

func countLinesFiles(f *os.File, counts map[string]int, filenames map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if filenames[input.Text()] == "" {
			filenames[input.Text()] = f.Name()
		}
	}
	// NOTE : ignoring potential erros from input.Err()
}
