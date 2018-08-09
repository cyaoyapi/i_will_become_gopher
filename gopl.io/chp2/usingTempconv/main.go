// This program converts its numeric command-line argument
// to Celsius and Fahrenheit. It uses the package 'tempconv'

package main

import (
	"bufio"
	"fmt"
	"i_will_become_gopher/gopl.io/chp2/tempconv" // local path
	"os"
	"strconv"
)

func main() {
	if len(os.Args[1:]) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			printErrorOrResult(input.Text())
		}
	} else {
		for _, arg := range os.Args[1:] {
			printErrorOrResult(arg)
		}
	}
}

func printErrorOrResult(s string) {
	argFloat64, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error : %v", err)
		os.Exit(1)
	}
	f := tempconv.Fahrenheit(argFloat64)
	c := tempconv.Celsuis(argFloat64)
	fmt.Println(argFloat64, " equal ", tempconv.CToF(c), "and ", tempconv.FToC(f))
}
