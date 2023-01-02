package main

import (
	"fmt"
	"os"
)

// Modify any of the echo program to also print `os.Args[0]`, the name of the command that invoked it.

func main() {
	var args, sep string
	for i := 0; i < len(os.Args); i++ {
		args += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(args)
}
