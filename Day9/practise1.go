package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
Ramesh basic salary is entered through the keyboard. His dearness allowance is 40% of basic salary, and house rent
allowance is 20% of basic salary. Write a program to calculate his gross salary.
*/

func main() {
	var basicSalary int
	basicSalary, _ = strconv.Atoi(os.Args[1])

	dearnessAllowance := (40 * basicSalary) / 100
	houseRentAllowance := (20 * basicSalary) / 100

	grossSalary := basicSalary + dearnessAllowance + houseRentAllowance

	fmt.Printf("Basic Salary         : %d \n", basicSalary)
	fmt.Printf("Dearness Allowance   : %d \n", dearnessAllowance)
	fmt.Printf("House Rent Allowance : %d \n", houseRentAllowance)
	fmt.Println("-----------------------------")
	fmt.Printf("Gross Salary         : %d \n", grossSalary)
}
