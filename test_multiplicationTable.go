// test_multiplicationTable
// Challenge - Write a program that prints a multiplication table for numbers up to 12.

package main

import (
	"fmt"
	"time"
)

func main() {

	targetNumber := 12
	leftColumn := 1
	count := 1 // also represents starting number
	result := 0

	/////////print the header////////////////////////////////
	fmt.Printf("X	")
	for count <= targetNumber {
		fmt.Printf("%d	", count)
		count++
	}
	count = 1 // reset count after printing the top columns
	///////////////////////////////////////////////////////

	for leftColumn <= 12 {

		fmt.Printf("\n")
		fmt.Printf("%d	", leftColumn) // print the left column 1,2,3,4

		for count <= targetNumber {

			// do the maths
			result = count * leftColumn
			fmt.Printf("%d	", result)
			count++
			time.Sleep(10 * time.Millisecond)
		}

		result = 0
		count = 1
		leftColumn++
	}
	fmt.Printf("\n")

}
