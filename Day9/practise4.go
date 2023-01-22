package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
Temperature of a city in Fahrenheit degree is input through keyboard. Write a program to convert this temperature
into Centigrade degree.
*/

func main() {
	fmt.Print("Please enter temperature in Fahrenheit : ")
	inputs := bufio.NewScanner(os.Stdin)
	inputs.Scan()

	tempFahrenheit, _ := strconv.ParseFloat(inputs.Text(), 32)

	// Formula to convert Fahrenheit to Celsius C= (F-32) * 100 / 180

	var tempCelsius float32 = float32(((tempFahrenheit - 32) * 100) / 180)

	fmt.Printf("Temperature in Centigrade : %.2f \n", tempCelsius)
}
