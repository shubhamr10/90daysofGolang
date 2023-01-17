package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// ex2 programs takes the input from command lines for the URLs,
// check if the http:// is present in the URL and add it if not there.

func main() {
	for _, url := range os.Args[1:] {
		hasPrefix := strings.HasPrefix(url, "http://")
		if !hasPrefix {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error while fetching: %v\n", err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error while fetching: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("----------------------------------------------------")
	}
}
