package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
Two numbers are input through the keyboard into two location C and D. Write a program to interchange the content of C and D
*/

func main() {
	c, _ := strconv.Atoi(os.Args[1])
	d, _ := strconv.Atoi(os.Args[2])
	swapC := c
	swapD := d
	fmt.Printf("Before swapping C: %d, D:%d \n", c, d)
	e := c
	c = d
	d = e

	fmt.Printf("After swapping C: %d, D:%d \n", c, d)
	swapIntelligent(swapC, swapD)
}

func swapIntelligent(swapC, swapD int) {
	fmt.Printf("Before swapping intelligent C: %d, D:%d \n", swapC, swapD)
	swapC, swapD = swapD, swapC
	fmt.Print(swapC, swapD)
	fmt.Printf("After swapping intelligent C: %d, D:%d \n", swapC, swapD)
}
