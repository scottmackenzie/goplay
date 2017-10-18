// Slices and Arrays
package main

import (
	"fmt"
)

func printer(words [4]string) { // declare the size or it'll be a slice
	for _, word := range words {
		fmt.Printf("%s ", word)

	}
	fmt.Printf("\n")
	words[2] = "blue" // note how this is no changed as the array is copied
}

func slicePrinter(words2 []string) { // declare the size or it'll be a slice
	for _, w := range words2 {
		fmt.Printf("%s ", w)

	}
	fmt.Printf("\n")
	words2[2] = "blue" // note how this is changed as we are using a slice now
}

func printerClean(list []string) { // declare the size or it'll be a slice
	for _, l := range list {
		fmt.Printf("%s ", l)

	}
	fmt.Printf("\n")
}

func main() {
	words := [...]string{"the", "quick", "brown", "fox"}
	//brackets index into an array.  ... means unknown number of values
	fmt.Printf("%s\n", words[2]) // read 3rd item in the array

	// Arrays are fixed pieces of memory that CANNOT be passed around by reference
	// Arrays are passed around by COPYING
	printer(words)
	printer(words)

	// It's better to use slices instead of Arrays.
	// Slices are windows into Arrays
	sliceWords := []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
	fmt.Printf("The length of the sliceWords slice is: %d\n", len(sliceWords))
	slicePrinter(sliceWords)
	slicePrinter(sliceWords) // Note "blue" has been changed so will say so on the printout this time

	// You can also modify slices to make slices of slices:

	slicePrinter(sliceWords[0:4])               // first 4 (0 to 3)
	slicePrinter(sliceWords[5:len(sliceWords)]) // from 5 to the end
	slicePrinter(sliceWords[5:])                // same as above but simpler (note "blue" still)
	slicePrinter(sliceWords[:3])                // First 3 (from 0 to but not including 3)

	sliceWords[7] = "brown" // cleanup

	// How to "Make" a slice and add to it later.... use "Make" of course

	moreWords := make([]string, 4) // 4 initiatised strings of 0
	moreWords[0] = "the"
	moreWords[1] = "quick"
	moreWords[2] = "brown"
	moreWords[3] = "fox"
	printerClean(moreWords)

	// another way would be to set a capacity but define a maximum

	moreWords2 := make([]string, 0, 4)     // Set a capacity for the slice: 0 lenth now but 4 max
	moreWords2 = append(moreWords2, "the") // use append to add to the slice®
	moreWords2 = append(moreWords2, "quick")
	moreWords2 = append(moreWords2, "brown")
	moreWords2 = append(moreWords2, "fox :-)")
	printerClean(moreWords2)
	fmt.Printf("Slice moreWords2 length is: %d and it can store %d\n", len(moreWords2), cap(moreWords2))

	//lets extend the slice correctly

	moreWords2 = append(moreWords2, "jumps") // this will automatically extend the slice via a copy
	printerClean(moreWords2)
	fmt.Printf("Slice moreWords2 length is: %d and it can store %d\n", len(moreWords2), cap(moreWords2))

	//lets extend the slice incorrectly (expect a "panic: runtime error: index out of range" error)
	//moreWords2[9] = "crashes"

	// What if you want to copy from one slice to another

	moreWords2Copied := make([]string, 5)
	copy(moreWords2Copied, moreWords2) // new command "copy" - used to make a copy of data
	printerClean(moreWords2Copied)
	fmt.Printf("Copy the slices and change:\n\n")
	printerClean(moreWords2) // show you are changing the data in the new slice
	moreWords2Copied[2] = "blue"
	printerClean(moreWords2Copied)
}
