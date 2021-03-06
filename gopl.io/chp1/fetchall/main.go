// This program fetches URLs in parallel and reports their times and sizes

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start the goroutine
	}
	logfile, err := os.Create("log")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error : %v\n", err)
		os.Exit(1)
	}
	for range os.Args[1:] {
		io.WriteString(logfile, <-ch)
	}

	fmt.Fprintf(logfile, "\n%.2fs elapsed\n", time.Since(start).Seconds())
	logfile.Close()
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // Send the error's message to channel
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("While reading %s : %v", url, err) // Send the error's message to channel
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s\n", secs, nbytes, url)
}
