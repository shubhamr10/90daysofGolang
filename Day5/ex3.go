package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// ex3 modify the fetch program to print the HTTP status code, found in the resp.Status

func main() {
	for _, url := range os.Args[1:] {
		resp, _ := http.Get(url)
		b, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Fprintf(os.Stdout, "resp code: %v :: %s", resp.Status, b)
	}

}
