package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
If the marks obtained by a student in five different subjects are input through the keyboard, find out the aggregate
marks and percentage marks obtained by the student. Assume that the maximum marks that can be obtained by a student
in each subject is 100.
*/

func main() {
	var marks []int
	fmt.Println("Please enter your marks")

	inputs := bufio.NewScanner(os.Stdin)
	for i := 0; i < 5; i++ {
		fmt.Printf("Enter your marks for subject %d : ", i+1)
		inputs.Scan()

		mark, _ := strconv.Atoi(inputs.Text())
		marks = append(marks, mark)
	}

	var totalMarks int = 0
	for _, mark := range marks {
		totalMarks = totalMarks + mark
	}
	var percentageObtained = float32(totalMarks) / 500
	fmt.Printf("Total marks obtained : %d \n", totalMarks)
	fmt.Printf("Percentage obtained  : %.2f \n", percentageObtained*100)

}
