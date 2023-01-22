package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

/*
If a five-digit number is input through the keyboard, write a program to reverse the number.
*/

func main() {
	num, _ := strconv.Atoi(os.Args[1])
	var newNum int = 0

	for i := 4; i >= 0; i-- {
		tens := int(math.Pow(10, float64(i)))
		mulTens := int(math.Pow(10, float64(4-i)))
		newNum = newNum + num/tens*mulTens
		num = int(math.Mod(float64(num), float64(tens)))
	}
	fmt.Printf("\nnumber:%d \n", newNum)
}
