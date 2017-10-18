// computerSumAndComputeProduct.go
// Challenge (1):  Write a program that asks the user for a number n and gives them the possibility to choose between computing the sum and computing the product of 1,…,n.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func calculateSum(inbound int) {

	count := 1
	total := 0

	for input := inbound; count <= input; count++ {
		fmt.Printf("%d ", count)
		total = total + count
		if count != input {
			fmt.Printf("+ ")
		}

	}
	fmt.Printf("\nTotal SUM is: %d\n", total)
}

func calculateProduct(inbound int) {

	count := 1
	total := 1

	for input := inbound; count <= input; count++ {
		fmt.Printf("%d ", count)
		total = total * count
		if count != input {
			fmt.Printf("* ")
		}

	}
	fmt.Printf("\nTotal PRODUCT is: %d\n", total)
}

func main() {

	var stop bool

	for !stop {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Printf("Please enter an int: --> ")
		scanner.Scan()
		inputNumber, ok := strconv.Atoi(scanner.Text()) // number is in input

		if ok != nil {
			os.Exit(1)
		}

		fmt.Printf(`Please enter "product" or "sum"  --> `)
		scanner.Scan()
		inputWord := scanner.Text() // string

		if ok != nil {
			os.Exit(1)
		}

		if inputWord == "product" || inputWord == "Product" || inputWord == "p" {
			calculateProduct(inputNumber)
		} else if inputWord == "sum" || inputWord == "Sum" || inputWord == "s" {
			calculateSum(inputNumber)
		}
	}
}
