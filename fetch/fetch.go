//Fetch prints the content found at each specified url
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		//ex1.8modify fetch to add the prefix http:// to each argument URL if it is missing
		if strings.HasPrefix(url, "http://") {
			execfetch(url)
		} else {
			url = "http://" + url
			execfetch(url)
		}
	}
}

func execfetch(url string) {

	//http Get function is from the net/http package
	//http.Get makes an http request and if there is no error returns the result in the response struct resp

	resp, err := http.Get(url)

	//ping a non-existent url to execute this section of code below to display the error
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	//ex1.9 modify fetch to also print the http status code
	fmt.Printf("fetch: Response Status %v\n", resp.Status)
	//The body field of the response struct contains the server response as a readable stream
	//ioutil.ReadAll is a function from io/ioutil package that reads the entire response and stores it in b
	//b, err := ioutil.ReadAll(resp.Body)
	//ex1.7 use io.Copy(dst,src) instead of ioutil.ReadAll to copy the response body to os.Stdout
	//closing Body stream to avoid leaking resources
	b, err := io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	//Printf will write the response body to standard output or for ex 1.7 the bytes receives in response
	fmt.Printf("Bytes %v\n", b)
}
