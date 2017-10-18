// test_selectOddItemsInAList

// Challenge - Write a function that returns the elements on odd positions in a list.

package main

import (
	//"errors"
	"fmt"
	"time"
)

func returnOddPositionsOnAList(inboundSlice []string) {
	for i := 0; i < len(inboundSlice); i = i + 2 {
		fmt.Printf("%d: %s\n", i+1, inboundSlice[i])
	}
	return
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
	words = append(words, "Mouse!")

	returnOddPositionsOnAList(words)

	elapsed := time.Since(start)
	fmt.Printf("\n\nProgram took %s\n", elapsed)
}
