//Fetchall fetches URLs in parallel using goroutines and reports their times and sizes
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	//a channel is a communication mechanism that allows one goroutine to pass values of a specific type to another goroutine
	ch := make(chan string) //creating a bidirectional channel
	for _, url := range os.Args[1:] {
		//ex1.8modify fetch to add the prefix http:// to each argument URL if it is missing
		if strings.HasPrefix(url, "http://") {
			go execfetch(url, ch) //asynchronous call execfetch
		} else {
			url = "http://" + url
			go execfetch(url, ch)
		}
	}

	//this for loop receives the summary results for each execfetch call and prints those lines
	//Having main do all the printing ensures that the output from each goroutine is processed as a unit
	//with no danger of interleaving if two goroutines finish at the same time.
	for range os.Args[1:] {
		fmt.Println(<-ch) //receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func execfetch(url string, ch chan<- string) {
	start := time.Now()

	//http Get function is from the net/http package
	//http.Get makes an http request and if there is no error returns the result in the response struct resp

	resp, err := http.Get(url)

	//ping a non-existent url to execute this section of code below to display the error
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	//note io.Copy reads the body of the response and discards it by writing to ioutil.Discard output stream
	//io.Copy returns the byte count of the response and an error if any
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	//as each result arrives for the urls being fetched, execfetch returns a summary to the channel
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)

}
