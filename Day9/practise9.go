package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
If a four-digit number is entered through the keyboard, write a number to obtain the sum of the first and last digit of
this number.
*/

func main() {
	num, _ := strconv.Atoi(os.Args[1])

	firstDigit := num / 1000
	lastDigit := num % 10

	fmt.Printf("Sum of first and last digits will be: %d", firstDigit+lastDigit)

}
