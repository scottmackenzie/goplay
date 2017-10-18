// typesAndKeywords
package main

import (
	"fmt"
	"time"
)

type summableSlice []int

// Types can be structures
// They can also be, in a way, aliases for built in types

func (s summableSlice) sum() int {
	sum := 0
	for _, i := range s {
		sum += i
	}
	return sum
}

func main() {
	start := time.Now()

	var s summableSlice = summableSlice{1, 2, 3, 4, 5, 6, 7, 8, 10}

	fmt.Printf("The sum is: %d\n", s.sum())

	elapsed := time.Since(start)
	fmt.Printf("\n\nProgram took %s\n", elapsed)
}
