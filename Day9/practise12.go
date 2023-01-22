package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
If the total selling price of 15 items and total profit earned on them is input through the keyboard, write a program
to find the cost price of one item.
*/

func main() {
	inputs := bufio.NewScanner(os.Stdin)

	fmt.Printf("Enter the selling price of 15 items: ")
	inputs.Scan()
	sellingPrice, _ := strconv.ParseFloat(inputs.Text(), 32)

	fmt.Printf("Enter the profit earned on 15 items: ")
	inputs.Scan()
	totalProfit, _ := strconv.ParseFloat(inputs.Text(), 32)

	totalCostPrice := sellingPrice - totalProfit

	fmt.Printf("Cost price of each item will be : %.2f \n", totalCostPrice/15)
}
