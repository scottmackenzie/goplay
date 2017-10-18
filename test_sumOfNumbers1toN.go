// test_sumOfNumbers1toN
// Challenge (1):  Write a program that asks the user for a number n and prints the sum of the numbers 1 to n
// Challenge (2):  Modify the previous program such that only multiples of three or five are considered in the sum, e.g. 3, 5, 6, 9, 10, 12, 15 for n=17

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func multipleOf3or5(input int) (result bool) {
	result = input%3 == 0 || input%5 == 0
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please enter an int: --> ")
	for scanner.Scan() {
		userInput, ok := strconv.Atoi(scanner.Text()) // number is in input
		if ok != nil {
			os.Exit(1)
		}

		// Sum of numbers means 1 + x + x + x through to n

		count := 1
		total := 0

		for input := userInput; count <= input; count++ {
			if multipleOf3or5(count) {
				fmt.Printf("%d ", count)
				total = total + count
			} else {
			}
		}
		fmt.Printf("\nTotal is: %d\n", total)
		fmt.Printf("Please enter an int: --> ")
	}
}
