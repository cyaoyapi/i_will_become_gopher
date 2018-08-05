package main 

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println("Liste of command-line arguments")
	for key,val := range os.Args[1:]{
		fmt.Println(key+1," : ",val)
	}
}