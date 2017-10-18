// test_askNameAndGreet
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please enter your name: \n-->")
	for scanner.Scan() {
		input := scanner.Text()
		switch {
		case input == "":
			os.Exit(1)
		case input == "bob" || input == "Bob" || input == "alice" || input == "Alice":
			fmt.Printf("\nHello %s\n", input)
		default:
			fmt.Printf("Thanks\n\n")
		}
		fmt.Printf("Please enter your name: \n-->")
	}
}
