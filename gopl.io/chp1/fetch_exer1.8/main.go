// This program prints the content found at a URL
// Add prefix "http://" to the given url if is missing

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error : %v\n", err)
			os.Exit(1)
		}
		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error : %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Response status : ", resp.Status)
		fmt.Println("==============================================")
		fmt.Printf("%s\n", body)

	}
}
