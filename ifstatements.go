// if statements
package main

import (
	"fmt"
	"os"
)

func main() {
	if numberBytes, theError := fmt.Println("Hello, World!"); theError != nil {
		// by default theError is nil if nothing has gone wrong
		// note numberBytes is only defined IN THE SCOPE of this if statement, not outside
		os.Exit(1)
	} else { // this else ensures the numberBytes included in flow
		fmt.Printf("Printed %d characters\n", numberBytes)
	}
	// alterntive way to do this would be

	numberBytes, theError := fmt.Println("Hello, World (round two)!")
	if theError != nil {
		os.Exit(1)
	} else {
		fmt.Printf("Printed %d characters\n", numberBytes)
	}

	// and more control
	numberBytes, theError = fmt.Println("Hello, World (round three)!")
	if theError != nil {
		os.Exit(1)
	} else if numberBytes > 20 {
		fmt.Printf("Printed %d characters\n", numberBytes)
	}

	// Don't care about anything else, just tell me if an error happened

	if _, theError = fmt.Printf("Hello, World (round four)!"); theError != nil {
		os.Exit(1)
	}

	// note the ; in the statement above. This is rare in go, only used in "for" statement as well.

	fmt.Printf("\n")
}
