// This program prints the command-line 
package main 

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	fmt.Println("This is the command-line : ",strings.Join(os.Args[:], " "))
}