package main

import (
	"fmt"
	"os"
)

// Modify the echo program to print the index and value of each of its arguments.

func main() {
	for index, args := range os.Args {
		fmt.Println(fmt.Sprintf("At index: %d, there is a argument: %s", index, args))
	}
}
