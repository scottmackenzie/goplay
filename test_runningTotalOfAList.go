// test_selectOddItemsInAList

// Challenge - Write a function that computes the running total of a list.

package main

import (
	//"errors"
	"fmt"
	"os"
	"time"
)

func totalSizeOfList(inboundSlice []string) (total int, e error) {
	total, e = fmt.Printf("Slice items = %d\n", len(inboundSlice))
	return total, e
}

func printer(inboundSlice []string) {
	for i, word := range inboundSlice {
		fmt.Printf("%d, %s\n", i, word)
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
	words = append(words, "Mouse!")
	words = append(words, "Chicken")

	_, err := totalSizeOfList(words)
	if err == nil {
		printer(words)
	} else {
		os.Exit(1)
	}

	words = append(words, "more 1")
	words = append(words, "more 2")
	words = append(words, "more 3")
	words = append(words, "more 4")

	printer(words)

	elapsed := time.Since(start)
	fmt.Printf("\n\nProgram took %s\n", elapsed)
}
