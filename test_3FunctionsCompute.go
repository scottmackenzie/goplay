// test_3FunctionsCompute

// test - Write three functions that compute the sum of the numbers
// in a list: using a for-loop, a while-loop and recursion.

package main

import (
	"fmt"
	"time"
)

func forLoopCount(listOfNumbers []int) {
	total := 0
	for number, _ := range listOfNumbers {
		fmt.Printf("Number:")
	}
	fmt.Printf("Total is: %d", total)
}

func main() {
	start := time.Now()

	numberList := []int{1, 2, 3}
	forLoopCount(numberList)

	fmt.Println("Hello World!")

	elapsed := time.Since(start)
	fmt.Printf("\n\nProgram took %s\n", elapsed)
}
