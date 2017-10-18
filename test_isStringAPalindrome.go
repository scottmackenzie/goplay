// test_isStringAPalindrome
// Write a function that tests whether a string is a palindrome.

package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	str1         string = "Claire"
	str2         string = "Nin"
	stringIsNull        = errors.New("Sorry the string sent is null")
)

// Reverse words in a string (copied from google)
func reverse(word string) string { // declare the size or it'll be a slice
	// make the word all lower case for the test
	r := []rune(word)
	//fmt.Printf("\n%d", r)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func checkIfaPalindrome(input string) {
	if input == "" {
		fmt.Printf("Error")
	}

	wordLowercase := strings.ToLower(input)
	//fmt.Printf("Upper: %s   Lower: %s\n", input, wordLowercase)

	reversedStr := reverse(wordLowercase) // reverse the string

	if wordLowercase == reversedStr {
		fmt.Printf("Success:\n%s\n%s\n\n", input, reversedStr)
	} else {
		fmt.Printf("Fail:\n%s\n%s\n\n", input, reversedStr)
	}
}

func main() {
	start := time.Now()

	checkIfaPalindrome(str1)
	checkIfaPalindrome(str2)

	// timing
	elapsed := time.Since(start)
	fmt.Printf("\n\nProgram took %s\n", elapsed)
}
