// test_largestElementInAList

// Challenge - largest element in a list

package main

import (
	"fmt"
	"time"
)

var previousSize int
var biggestWord string

func main() {
	start := time.Now()

	// lets create a slice

	moreWords := make([]string, 10)
	moreWords[0] = "the"
	moreWords[1] = "quick"
	moreWords[2] = "brown"
	moreWords[3] = "fox"
	moreWords[4] = "jumps"
	moreWords[5] = "over"
	moreWords[6] = "the"
	moreWords[7] = "Lazy chicken"
	moreWords[8] = "dog"
	moreWords[9] = "Elephant!" // longest item in the slice

	biggestWordSoFar := moreWords[0] // set to first item in slice

	for _, word := range moreWords {
		fmt.Printf("Previous word: %s\n", word)
		if len(word) > len(biggestWordSoFar) { // if current word is bigger than the biggest word so far
			biggestWordSoFar = word
		}
		previousSize = len(word)
		//fmt.Printf("Biggest so far word: %s\n\n", biggestWordSoFar)
	}

	fmt.Printf("The biggest word is: %s (%d)\n", biggestWordSoFar, len(biggestWordSoFar))

	elapsed := time.Since(start)
	fmt.Printf("Program took %s\n", elapsed)
}
