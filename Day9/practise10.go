package main

import "fmt"

/*
In the town, the percentage of men is 52. The percentage of total literacy is 48. If total percentage of literate men
is 35 of the total population. Write a program to find the total number of illiterate men and women if the population
of the town is 80,000
*/

func main() {
	totalPopulation := 80000
	percentageOfMen := 52
	percentageOfTotalLiteracy := 48
	percentageOfLiterateMen := 35

	totalLiteratePopulation := (percentageOfTotalLiteracy * totalPopulation) / 100

	totalMen := (percentageOfMen * totalPopulation) / 100
	totalLiterateMen := (percentageOfLiterateMen * totalMen) / 100

	totalWomen := totalPopulation - totalMen
	totalLiterateWomen := totalLiteratePopulation - totalLiterateMen

	totalIlliterateMen := totalMen - totalLiterateMen
	totalIlliterateWomen := totalWomen - totalLiterateWomen

	fmt.Printf("Total population       : %d \n", totalPopulation)
	fmt.Printf("Total Men              : %d \n", totalMen)
	fmt.Printf("Total Women            : %d \n", totalWomen)
	fmt.Printf("Total Literate Men     : %d \n", totalLiterateMen)
	fmt.Printf("Total Literate Women   : %d \n", totalLiterateWomen)
	fmt.Printf("Total Illiterate Men   : %d \n", totalIlliterateMen)
	fmt.Printf("Total Illiterate Women : %d \n", totalIlliterateWomen)
}
