// This program prints its arguments : version 2

package main 

import (
	"fmt"
	"os"
)

func main(){
	s, sep := "",""
	for _, arg := range os.Args[1:]{
		s += sep + arg
		sep = " "
	}
	fmt.Println("The list of the given arguments : [",s,"]")

}