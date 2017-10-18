// switches
package main

import (
	"fmt"
	"os"
)

func main() {
	n, err := fmt.Printf("Hello World!\n")

	// note switch statements do not fallthrough unless defined to

	switch {
	case err != nil:
		os.Exit(1)
	case n == 0: // you can add & if you like to make the switch more complex
		fmt.Printf("No bytes output\n")
	case n != 14:
		fmt.Printf("Wrong number of characters: %d\n", n)
		fallthrough // rare but can be used
	default:
		fmt.Printf("OK!\n")
	}

	// lets try another flow

	atoz := "The quick brown fox jumps over the lazy dog"
	vowels := 0
	consonants := 0
	z := 0

	for _, r := range atoz {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			vowels += 1
		case 'z':
			z += 1
			fallthrough
		default:
			consonants += 1
		}
	}
	fmt.Printf("Vowles: %d, Consonants %d and z: %d\n", vowels, consonants, z)
}
