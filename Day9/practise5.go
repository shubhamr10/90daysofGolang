package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

/*
The length and breadth of a rectangle and radius of a circle are input through the keyboard. Write a program to
calculate the area& perimeter of the rectangle, and the area & circumference of the circle.
*/

func main() {
	fmt.Println("Enter the length of the rectangle :")
	inputs := bufio.NewScanner(os.Stdin)
	inputs.Scan()
	lengthRectange, _ := strconv.ParseFloat(inputs.Text(), 32)

	fmt.Println("Enter the breadth of the rectangle :")
	inputs.Scan()
	breadthRectange, _ := strconv.ParseFloat(inputs.Text(), 32)

	fmt.Println("Enter the radius of the circle :")
	inputs.Scan()
	radiusCircle, _ := strconv.ParseFloat(inputs.Text(), 32)

	fmt.Printf("The area of a rectange of length : %.2f and breadth : %.2f is : %.2f \n", lengthRectange, breadthRectange, lengthRectange*breadthRectange)
	fmt.Printf("The perimeter of a rectange of length : %.2f and breadth : %.2f is : %.2f \n", lengthRectange, breadthRectange, 2*(lengthRectange+breadthRectange))

	fmt.Printf("The area of circle with radius : %.2f is : %.2f \n", radiusCircle, math.Pi*math.Pow(radiusCircle, 2))
	fmt.Printf("The circumference of a circle with radius : %.2f is %.2f \n", radiusCircle, math.Pi*2*radiusCircle)

}
