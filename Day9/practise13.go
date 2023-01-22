package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
If a five-digit number is input through the keyboard, write a program to print a new number by adding one to each of its
digits. For example, if the number is eis 12391 then output should be 23402
*/

func main() {
	fmt.Printf("Enter a five digit number :")
	inputs := bufio.NewScanner(os.Stdin)

	inputs.Scan()
	num, _ := strconv.Atoi(inputs.Text())

	// instead of looping through each digit, we will just add 11111 to the number
	num = num + 11111

	fmt.Printf("Add 1 to each digit: %d \n", num)
}
