package main

import (
	"fmt"
	"os"
)

func main() {
	var args, separate string
	for i := 1; i < len(os.Args); i++ {
		args += separate + os.Args[i]
		separate = " "
	}
	fmt.Println(args)
}
