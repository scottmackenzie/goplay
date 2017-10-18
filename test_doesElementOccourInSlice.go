// test_doesElementOccourInSlice
// Write a function that checks whether an element occurs in a list.

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please enter a URL --> ")
	for scanner.Scan() {
		userInput := scanner.Text() // number is in input

		// Rather than using a slice, lets use a map

		moreWords := map[string]bool{
			"www.google.com":   true,
			"www.yahoo.co.uk":  true,
			"www.edapt.org.uk": true,
			"news.bbc.co.uk":   true,
			"www.cnn.com":      true,
			"www.gov.uk":       true,
			"www.truphone.com": true,
			"www.gsma.org":     true,
		} // true if it's in the map

		if moreWords[userInput] {
			fmt.Printf("URL found: %s\n", userInput)
		} else {
			fmt.Printf("URL not found %s\n", userInput)
		}

		elapsed := time.Since(start)
		fmt.Printf("Program took %s\n", elapsed)

		fmt.Printf("\n\nPlease enter a URL --> ")

	}
}
