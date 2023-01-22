package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
A cashier has currency notes of denomination 10, 50 and 100. If the amount to be withdrawn is input through keyboard
in hundreds, find the total number of currency notes of each denomination the cashier will have to give to the withdrawal.
*/

func main() {
	withdrawalAmount, _ := strconv.Atoi(os.Args[1])
	var (
		hundreds = 0
		tens     = 0
		fifties  = 0
	)
	hundreds = withdrawalAmount / 100
	withdrawalAmount = withdrawalAmount % 100
	if withdrawalAmount > 0 {
		fifties = withdrawalAmount / 50
		withdrawalAmount = withdrawalAmount % 50
	}
	if withdrawalAmount > 0 {
		tens = withdrawalAmount / 10
		withdrawalAmount = withdrawalAmount % 10
	}

	fmt.Printf("Hundreds :%d, fifties : %d, tens : %d \n", hundreds, fifties, tens)
}
