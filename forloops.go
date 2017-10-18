// forloops
package main

import (
	"fmt"
)

func main() {
	counter := 0

	// less than 10
	for counter < 10 {
		fmt.Printf("Hello, World!\n")
		counter += 1
	}

	// another way to write this would be (semi colon seperate the 3 clauses)
	for counter2 := 0; counter2 < 10; counter2 += 1 {
		fmt.Printf("Counter: %d\n", counter2)
	}

	// another way to write this would be smaller:
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %d\n", i)
	}

	// Many variables (note the two variables, and their initialisation seperated by commas
	for i, j := 0, 1; i < 10; i, j = i+1, j*2 {
		fmt.Printf("i: %d, j: %d\n", i, j)
	}

	// boolean based while loop
	var stop bool
	counter3 := 0

	for !stop {
		fmt.Printf("%d: Hello, World!\n", counter3)
		counter3++
		if counter3 == 10 {
			stop = true
		}
		fmt.Printf("Bool is: %t  :   ", stop)
	}
}
