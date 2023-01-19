package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/count", counter)
	http.HandleFunc("/", handler2)

	log.Fatal(http.ListenAndServe("localhost:8100", nil))
}

// handler2 echoes the Path component of the requested URL
func handler2(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.path = %q %d \n", r.URL.Path, count)
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d \n", count)
	mu.Unlock()
}
