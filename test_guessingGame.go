// test_guessingGame

// Challenge - Write a guessing game where the user has to guess a secret number. After every guess the program tells the user whether their number was too large or too small. At the end the number of tries needed should be printed. I counts only as one try if they input the same number multiple times consecutively.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var stop bool
	sliceInput := []int{0}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please enter the search number: --> ")
	scanner.Scan()
	inputNumber, ok := strconv.Atoi(scanner.Text()) // number is in input
	fmt.Printf("%d\n", inputNumber)

	if ok != nil {
		os.Exit(1)
	}
	for !stop {

		fmt.Printf("Guess the number... \n")
		scanner.Scan()
		inputGuess, ok := strconv.Atoi(scanner.Text()) // number is in input
		// write each input to the sliceInput
		if ok != nil {
			os.Exit(1)
		}
		//fmt.Printf("Input guess = %d and sliceInput = %d\n", inputGuess, sliceInput[len(sliceInput)-1])
		if inputGuess != sliceInput[len(sliceInput)-1] {
			sliceInput = append(sliceInput, inputGuess)
		} else {
			fmt.Printf("Entry same as last time\n")
		}
		if inputGuess < inputNumber {
			fmt.Printf("Higher...\n")
		} else if inputGuess > inputNumber {
			fmt.Printf("Lower...\n")
		} else if inputGuess == inputNumber {
			fmt.Printf("\nWINNER! (in %d tries)...\n", len(sliceInput)-1)
			break
		}
	}

}
