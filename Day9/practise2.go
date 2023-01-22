package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
The distance between two cities (in km.) is input through the keyboard. Write a program to convert and print this
distance in meters, feet, inches and centimeters.
*/

func main() {
	distanceInKM, _ := strconv.Atoi(os.Args[1])

	distanceInMeters := float32(distanceInKM * 1000)
	distanceInCentimeters := distanceInMeters * 100
	distanceInFeet := distanceInMeters * 3.28084
	distanceInInches := distanceInMeters * 39.37008

	fmt.Printf("Distance in Kilometers  : %d \n", distanceInKM)
	fmt.Printf("Distance in Centimeters : %.2f \n", distanceInCentimeters)
	fmt.Printf("Distance in Meters      : %.2f \n", distanceInMeters)
	fmt.Printf("Distance in Inches      : %.2f \n", distanceInInches)
	fmt.Printf("Disntance in Feet       : %.2f \n", distanceInFeet)

}
