// test_reverseAListInPlace

// Challenge - Write function that reverses a list, preferably in place.

package main

import (
	"fmt"
	"time"
)

func reverse(words []string) { // declare the size or it'll be a slice

	// read how long the array is and print it out backwards

	sliceSize := len(words)
	//fmt.Printf("Slice size (reverse) is: %d\n", sliceSize)

	for i := sliceSize; i > 0; i-- {
		//fmt.Printf("i is:  %d \n", i)
		fmt.Printf("%s ", words[i-1])
	}
}

func main() {

	var words []string // create slice of unknown size
	start := time.Now()

	// lets create a slice (list)

	words = append(words, "the")
	words = append(words, "quick")
	words = append(words, "brown")
	words = append(words, "fox")
	words = append(words, "jumps")
	words = append(words, "over")
	words = append(words, "the")
	words = append(words, "lazy")
	words = append(words, "dog")
	words = append(words, "Elephant!")

	//fmt.Printf("Slice size (main) is: %d\n", len(words))
	for _, word := range words {
		fmt.Printf("%v ", word)
	}

	fmt.Printf("\n")
	reverse(words)

	elapsed := time.Since(start)
	fmt.Printf("\n\nProgram took %s\n", elapsed)
}
